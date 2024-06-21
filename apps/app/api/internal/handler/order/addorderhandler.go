package order

import (
	"net/http"

	"github.com/billbliu/lebron/apps/app/api/internal/logic/order"
	"github.com/billbliu/lebron/apps/app/api/internal/svc"
	"github.com/billbliu/lebron/apps/app/api/internal/types"
	"github.com/billbliu/lebron/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// add order
func AddOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderAddReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := order.NewAddOrderLogic(r.Context(), svcCtx)
		resp, err := l.AddOrder(&req)
		result.HttpResult(r, w, resp, err)
	}
}
