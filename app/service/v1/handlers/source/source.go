package source

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hramov/xreport/business/platform/database"
	"github.com/hramov/xreport/business/platform/web"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	shutdownCh chan os.Signal
	db         *sqlx.DB
	log        *log.Logger
}

type Source struct {
	Id        string    `db:"id" json:"id"`
	DriverId  int       `db:"driver_id" json:"driver_id"`
	Driver    string    `db:"driver" json:"driver"`
	Title     string    `db:"title" json:"title"`
	Host      string    `db:"host" json:"host"`
	Port      int       `db:"port" json:"port"`
	Username  string    `db:"username" json:"user"`
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
	app.Handle("POST", "/source/check", handler.check)
	app.Handle("POST", "/source/", handler.create)
	app.Handle("PUT", "/source/", handler.update)
}

func (h *Handler) list(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	s := []Source{}
	err := h.db.SelectContext(ctx, &s, "select * from source")
	return s, err
}

func (h *Handler) check(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	var s Source
	err := web.Decode(r, &s)
	if err != nil {
		return nil, web.AppError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	_, err = database.New(database.Db(s.Driver), s.Host, s.Port, s.Username, s.Password, s.Database)
	if err != nil {
		return nil, web.AppError{
			Code:    http.StatusForbidden,
			Message: err,
		}
	}

	return true, nil
}

func (h *Handler) create(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	var s Source
	err := web.Decode(r, &s)
	if err != nil {
		return nil, err
	}

	sql := `insert into source (driver_id, title, host, port, username, password, database, created_at) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`
	row := h.db.QueryRowContext(ctx, sql, s.DriverId, s.Title, s.Host, s.Port, s.Username, s.Password, s.Database, s.CreatedAt)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var id int
	row.Scan(&id)

	return id, nil
}

func (h *Handler) update(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
	var s Source
	err := web.Decode(r, &s)
	if err != nil {
		return nil, err
	}

	sql := `update source set driver_id = $1, title = $2, host = $3, port = $4, username = $5, password = $6, database = $7, created_at = $8) returning id`
	row := h.db.QueryRowContext(ctx, sql, s.DriverId, s.Title, s.Host, s.Port, s.Username, s.Password, s.Database, s.CreatedAt)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var id int
	row.Scan(&id)

	return id, nil
}
