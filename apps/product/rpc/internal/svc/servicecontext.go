package svc

import (
	"time"

	"github.com/billbliu/lebron/apps/product/rpc/internal/config"
	"github.com/billbliu/lebron/apps/product/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
)

const localCacheExpire = time.Duration(time.Second * 60)

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	OperationModel model.ProductOperationModel
	SingleGroup    singleflight.Group
	// 业务缓存Redis客户端
	BizRedis   *redis.Redis
	LocalCache *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:         c,
		ProductModel:   model.NewProductModel(conn, c.CacheRedis),
		CategoryModel:  model.NewCategoryModel(conn, c.CacheRedis),
		OperationModel: model.NewProductOperationModel(conn, c.CacheRedis),
		BizRedis:       redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		LocalCache:     localCache,
	}
}
