package v1

import (
	"database/sql"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/hramov/xreport/app/service/v1/handlers/source"
	"log"
)

func New(shutdownCh chan struct{}, db *sql.DB, log *log.Logger) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	source.New(shutdownCh, mux, db, log)

	return mux
}
