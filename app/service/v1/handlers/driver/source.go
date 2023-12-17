package driver

import (
	"context"
	"github.com/hramov/xreport/business/platform/web"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	shutdownCh chan os.Signal
	db         *sqlx.DB
	log        *log.Logger
}

type Driver struct {
	Id        string    `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Code      string    `db:"code" json:"code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func New(shutdownCh chan os.Signal, app *web.App, db *sqlx.DB, log *log.Logger) {
	handler := &Handler{
		shutdownCh: shutdownCh,
		db:         db,
		log:        log,
	}

	app.Handle("GET", "/driver/", handler.list)
}

func (h *Handler) list(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	s := []Driver{}
	err := h.db.SelectContext(ctx, &s, "select * from driver")
	return s, err
}
