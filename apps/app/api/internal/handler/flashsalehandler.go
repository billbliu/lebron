package handler

import (
	"net/http"

	"github.com/billbliu/lebron/apps/app/api/internal/logic"
	"github.com/billbliu/lebron/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 限时抢购
func FlashSaleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFlashSaleLogic(r.Context(), svcCtx)
		resp, err := l.FlashSale()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
