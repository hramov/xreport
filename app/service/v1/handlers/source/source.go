package source

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

type Source struct {
	Id        string    `db:"id" json:"id"`
	DriverId  int       `db:"driver_id" json:"driver_id"`
	Title     string    `db:"title" json:"title"`
	Host      string    `db:"host" json:"host"`
	Port      string    `db:"port" json:"port"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	Database  string    `db:"database" json:"database"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func New(shutdownCh chan os.Signal, app *web.App, db *sqlx.DB, log *log.Logger) {
	handler := &Handler{
		shutdownCh: shutdownCh,
		db:         db,
		log:        log,
	}

	app.Handle("GET", "/source/", handler.list)
}

func (h *Handler) list(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	s := []Source{}
	err := h.db.SelectContext(ctx, &s, "select * from source")
	return s, err
}
