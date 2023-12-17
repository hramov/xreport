package middleware

import (
	"context"
	"fmt"
	"github.com/hramov/xreport/business/platform/web"
	"log"
	"net/http"
)

func Panic(log *log.Logger) web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (res any, err error) {
			defer func() {
				recErr := recover()
				if recErr != nil {
					log.Println("panic:", recErr)
					res = nil
					err = fmt.Errorf("internal server error")
				}
			}()
			res, err = h(ctx, w, r)
			return
		}
	}
}
