package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/billbliu/lebron/apps/product/rpc/internal/model"
	"github.com/billbliu/lebron/apps/product/rpc/internal/svc"
	"github.com/billbliu/lebron/apps/product/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	defaultPageSize = 10
	defaultLimit    = 300
	expireTime      = 3600 * 24 * 3
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *product.ProductListRequest) (*product.ProductListResponse, error) {
	// 判断分类id是否存在
	_, err := l.svcCtx.CategoryModel.FindOne(l.ctx, uint64(in.CategoryId))
	if err == model.ErrNotFound {
		return nil, status.Error(codes.NotFound, "category not found")
	}

	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}

	if in.Ps == 0 {
		in.Ps = defaultPageSize
	}

	var (
		isCache, isEnd   bool  // 是否是缓存, 是否是最后一条数据
		lastID, lastTime int64 // 本页最后一条记录的ID和创建时间
		firstPage        []*product.ProductItem
		products         []*model.Product
	)

	// 查出缓存中的pids
	pids, _ := l.cacheProductList(l.ctx, in.CategoryId, in.Cursor, int64(in.Ps))
	if len(pids) == int(in.Ps) { // 能从缓存中查到请求的所有id
		isCache = true
		// 最后一条数据为-1, 已经加载到最后一页了
		if pids[len(pids)-1] == 0 {
			isEnd = true
		}

		// 根据缓存查出的id列表在mysql数据库中查出所有商品
		products, err := l.productsByIds(l.ctx, pids)
		if err != nil {
			return nil, err
		}

		for _, p := range products {
			firstPage = append(firstPage, &product.ProductItem{
				ProductId:  int64(p.Id),
				Name:       p.Name,
				CreateTime: p.CreateTime.Unix(),
			})
		}
	} else { // 不能从缓存中查到请求的所有id, 直接查mysql数据库
		var (
			err   error
			ctime = time.Unix(in.Cursor, 0).Format("2006-01-02 15:04:05")
		)

		// 根据过滤条件在数据库中查出所有商品
		products, err := l.svcCtx.ProductModel.CategoryProducts(l.ctx, ctime, int64(in.CategoryId), defaultLimit)
		if err != nil {
			return nil, err
		}

		var firstPageProducts []*model.Product
		if len(products) > int(in.Ps) {
			firstPageProducts = products[:int(in.Ps)]
			isEnd = true
		} else {
			firstPageProducts = products
			isEnd = true
		}

		for _, p := range firstPageProducts {
			firstPage = append(firstPage, &product.ProductItem{
				ProductId:  int64(p.Id),
				Name:       p.Name,
				CreateTime: p.CreateTime.Unix(),
			})
		}
	}

	if len(firstPage) > 0 {
		pageLast := firstPage[len(firstPage)-1]
		lastID = pageLast.ProductId
		lastTime = pageLast.CreateTime

		if lastTime < 0 {
			lastTime = 0
		}

		for k, p := range firstPage {
			if p.CreateTime == in.Cursor && p.ProductId == in.ProductId {
				firstPage = firstPage[:k]
				break
			}
		}
	}

	ret := &product.ProductListResponse{
		IsEnd:     isEnd,
		Timestamp: lastTime,
		ProductId: lastID,
		Products:  firstPage,
	}

	// 没有缓存则缓存起来
	if !isCache {
		threading.GoSafe(func() {
			if len(products) < defaultLimit && len(products) > 0 {
				endTime, _ := time.Parse("2006-01-02 15:04:05", "0000-00-00 00:00:00")
				products = append(products, &model.Product{Id: 0, CreateTime: endTime})
			}

			_ = l.addCacheProductList(context.Background(), products)
		})
	}

	return ret, nil
}

func (l *ProductListLogic) productsByIds(ctx context.Context, pids []int64) ([]*model.Product, error) {
	products, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pids {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer[*model.Product], cancel func(error)) {
		pid := item.(int64)
		p, err := l.svcCtx.ProductModel.FindOne(ctx, uint64(pid))
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(p)
	}, func(pipe <-chan *model.Product, writer mr.Writer[[]*model.Product], cancel func(error)) {
		var products []*model.Product
		for item := range pipe {
			products = append(products, item)
		}
		writer.Write(products)
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (l *ProductListLogic) cacheProductList(ctx context.Context, cid int32, cursor, ps int64) ([]int64, error) {
	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, l.categoryKey(cid), cursor, 0, 0, int(ps))
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, pair := range pairs {
		id, _ := strconv.ParseInt(pair.Key, 10, 64)
		ids = append(ids, id)
	}

	return ids, nil
}

func (l *ProductListLogic) addCacheProductList(ctx context.Context, products []*model.Product) error {
	if len(products) == 0 {
		return nil
	}

	for _, p := range products {
		score := p.CreateTime.Unix()
		if score < 0 {
			score = 0
		}

		if _, err := l.svcCtx.BizRedis.ZaddCtx(ctx, l.categoryKey(int32(p.Cateid)), score, strconv.Itoa(int(p.Id))); err != nil {
			return err
		}
	}
	return l.svcCtx.BizRedis.ExpireCtx(ctx, l.categoryKey(int32(products[0].Cateid)), expireTime)
}

func (l *ProductListLogic) categoryKey(cid int32) string {
	return fmt.Sprintf("category:%d", cid)
}
