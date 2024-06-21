package user

import (
	"net/http"

	"github.com/billbliu/lebron/apps/app/api/internal/logic/user"
	"github.com/billbliu/lebron/apps/app/api/internal/svc"
	"github.com/billbliu/lebron/apps/app/api/internal/types"
	"github.com/billbliu/lebron/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// edit user receiveAddress
func EditReceiveAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReceiveAddressEditReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewEditReceiveAddressLogic(r.Context(), svcCtx)
		resp, err := l.EditReceiveAddress(&req)
		result.HttpResult(r, w, resp, err)
	}
}
