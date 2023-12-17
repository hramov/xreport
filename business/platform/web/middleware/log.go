package middleware

import (
	"context"
	"github.com/hramov/xreport/business/platform/web"
	"log"
	"net/http"
)

func Log(log *log.Logger) web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
			reqId, ok := ctx.Value("req-id").(string)
			if !ok {
				reqId = "unknown"
			}

			log.Println("START : ", reqId, r.Method, r.URL.Path)
			res, err := h(ctx, w, r)
			log.Println("END   : ", reqId, r.Method, r.URL.Path, err)
			return res, err
		}
	}
}
