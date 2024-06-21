package user

import (
	"net/http"

	"github.com/billbliu/lebron/apps/app/api/internal/logic/user"
	"github.com/billbliu/lebron/apps/app/api/internal/svc"
	"github.com/billbliu/lebron/apps/app/api/internal/types"
	"github.com/billbliu/lebron/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// del user receiveAddress list
func DelReceiveAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReceiveAddressDelReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewDelReceiveAddressLogic(r.Context(), svcCtx)
		resp, err := l.DelReceiveAddress(&req)
		result.HttpResult(r, w, resp, err)
	}
}
