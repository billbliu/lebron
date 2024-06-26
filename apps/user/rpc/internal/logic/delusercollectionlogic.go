package logic

import (
	"context"

	"github.com/billbliu/lebron/apps/user/rpc/internal/model"
	"github.com/billbliu/lebron/apps/user/rpc/internal/svc"
	"github.com/billbliu/lebron/apps/user/rpc/user"
	"github.com/billbliu/lebron/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserCollectionLogic {
	return &DelUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收藏
func (l *DelUserCollectionLogic) DelUserCollection(in *user.UserCollectionDelReq) (*user.UserCollectionDelRes, error) {
	_, err := l.svcCtx.UserCollectionModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "该商品没有被收藏")
		}
		return nil, err
	}
	dbCollection := new(model.UserCollection)
	dbCollection.Id = uint64(in.Id)
	dbCollection.IsDelete = 1
	err = l.svcCtx.UserCollectionModel.UpdateIsDelete(l.ctx, dbCollection)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DelUserCollection Database Exception : %+v , err: %v", dbCollection, err)
	}
	return &user.UserCollectionDelRes{}, nil
}
