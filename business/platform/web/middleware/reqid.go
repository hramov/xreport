package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/hramov/xreport/business/platform/web"
	"net/http"
)

func ReqId() web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
			ctx = context.WithValue(ctx, "req-id", uuid.New().String())
			return h(ctx, w, r)
		}
	}
}
