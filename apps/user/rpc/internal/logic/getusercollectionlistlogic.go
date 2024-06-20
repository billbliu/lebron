package logic

import (
	"context"
	"fmt"

	"github.com/billbliu/lebron/apps/user/rpc/internal/model"
	"github.com/billbliu/lebron/apps/user/rpc/internal/svc"
	"github.com/billbliu/lebron/apps/user/rpc/user"
	"github.com/billbliu/lebron/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏列表
func (l *GetUserCollectionListLogic) GetUserCollectionList(in *user.UserCollectionListReq) (*user.UserCollectionListRes, error) {
	collectionList, err := l.svcCtx.UserCollectionModel.FindAllByUid(l.ctx, in.Uid)
	fmt.Println("==============collectionList", collectionList)
	fmt.Println("============== in.Uid", in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed  get user's Collection list err : %v , in :%+v", err, in)
	}
	var resp []int64
	for _, collections := range collectionList {
		//_ = copier.Copy(&resp, collections.ProductId)
		resp = append(resp, int64(collections.ProductId))
	}
	fmt.Println("============== resp", resp)
	return &user.UserCollectionListRes{
		ProductId: resp,
	}, nil
}
