package user

import (
	"context"
	"encoding/json"

	"github.com/billbliu/lebron/apps/app/api/internal/svc"
	"github.com/billbliu/lebron/apps/app/api/internal/types"
	"github.com/billbliu/lebron/apps/user/rpc/user"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// add user receiveAddress
func NewAddReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReceiveAddressLogic {
	return &AddReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReceiveAddressLogic) AddReceiveAddress(req *types.UserReceiveAddressAddReq) (resp *types.UserReceiveAddressAddRes, err error) {
	//get jwt token uid
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	var addRpcReq user.UserReceiveAddressAddReq
	addRpcReq.Uid = uid
	errCopy := copier.Copy(&addRpcReq, req)
	if errCopy != nil {
		return nil, errCopy
	}
	_, err = l.svcCtx.UserRPC.AddUserReceiveAddress(l.ctx, &addRpcReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
