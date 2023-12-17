package middleware

import (
	"context"
	"github.com/hramov/xreport/business/platform/web"
	"net/http"
)

// Cors sets the response headers needed for Cross-Origin Resource Sharing
func Cors(origin string) web.Middleware {
	return func(handler web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")

			return handler(ctx, w, r)
		}
	}
}
