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

type AddUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserReceiveAddressLogic {
	return &AddUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加收获地址
func (l *AddUserReceiveAddressLogic) AddUserReceiveAddress(in *user.UserReceiveAddressAddReq) (*user.UserReceiveAddressAddRes, error) {
	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Uid = uint64(in.GetUid())
	dbAddress.Name = in.GetName()
	dbAddress.Phone = in.GetPhone()
	dbAddress.Province = in.GetProvince()
	dbAddress.City = in.GetCity()
	dbAddress.IsDefault = uint64(in.GetIsDefault())
	dbAddress.PostCode = in.GetPostCode()
	dbAddress.Region = in.GetRegion()
	dbAddress.DetailAddress = in.GetDetailAddress()
	_, err := l.svcCtx.UserReceiveAddressModel.Insert(l.ctx, dbAddress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUserReceiveAddress Database Exception : %+v , err: %v", dbAddress, err)
	}
	return &user.UserReceiveAddressAddRes{}, nil
}
