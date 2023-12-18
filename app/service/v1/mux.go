package v1

import (
	"log"
	"os"

	"github.com/hramov/xreport/app/service/v1/handlers/driver"
	"github.com/hramov/xreport/app/service/v1/handlers/health"
	"github.com/hramov/xreport/app/service/v1/handlers/source"
	"github.com/hramov/xreport/business/platform/web"
	"github.com/hramov/xreport/business/platform/web/middleware"
	"github.com/jmoiron/sqlx"
)

func New(shutdownCh chan os.Signal, db *sqlx.DB, log *log.Logger) *web.App {
	app := web.New(shutdownCh, db, log)

	app.EnableCORS(middleware.Cors("*"))

	app.Use(middleware.Cors("*"), middleware.ReqId(), middleware.Panic(log))

	health.New(shutdownCh, app, db, log)
	source.New(shutdownCh, app, db, log)
	driver.New(shutdownCh, app, db, log)

	return app
}
