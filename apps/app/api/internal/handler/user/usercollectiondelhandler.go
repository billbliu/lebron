package user

import (
	"net/http"

	"github.com/billbliu/lebron/apps/app/api/internal/logic/user"
	"github.com/billbliu/lebron/apps/app/api/internal/svc"
	"github.com/billbliu/lebron/apps/app/api/internal/types"
	"github.com/billbliu/lebron/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// del user collection
func UserCollectionDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCollectionDelReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUserCollectionDelLogic(r.Context(), svcCtx)
		resp, err := l.UserCollectionDel(&req)
		result.HttpResult(r, w, resp, err)
	}
}
