package generate

import (
	"github.com/hramov/xreport/business/platform/web"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type Handler struct {
	shutdownCh chan os.Signal
	db         *sqlx.DB
	log        *log.Logger
}

type Request struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func New(shutdownCh chan os.Signal, app *web.App, db *sqlx.DB, log *log.Logger) {
	//handler := &Handler{
	//	shutdownCh: shutdownCh,
	//	db:         db,
	//	log:        log,
	//}

	//app.Handle("GET", "/", handler.list)
}
