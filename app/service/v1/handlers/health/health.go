package health

import (
	"context"
	"github.com/hramov/xreport/business/platform/web"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
)

type Handler struct {
	shutdownCh chan os.Signal
	db         *sqlx.DB
	log        *log.Logger
}

func New(shutdownCh chan os.Signal, app *web.App, db *sqlx.DB, log *log.Logger) {
	handler := &Handler{
		shutdownCh: shutdownCh,
		db:         db,
		log:        log,
	}

	app.Handle("GET", "/health", handler.health)
}

func (h *Handler) health(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	return "OK", nil
}
