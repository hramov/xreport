package source

import (
	"database/sql"
	"github.com/dimfeld/httptreemux/v5"
	"log"
	"net/http"
)

type Handler struct {
	shutdownCh chan struct{}
	db         *sql.DB
	log        *log.Logger
}

func New(shutdownCh chan struct{}, mux *httptreemux.ContextMux, db *sql.DB, log *log.Logger) {
	handler := &Handler{
		shutdownCh: shutdownCh,
		db:         db,
		log:        log,
	}
	group := mux.NewGroup("/source")
	handler.register(group)
}

func (h *Handler) register(mux *httptreemux.ContextGroup) {
	mux.GET("/", h.list)
	mux.GET("/{id}", h.get)
	mux.POST("/", h.create)
	mux.PUT("/{id}", h.update)
	mux.DELETE("/{id}", h.delete)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
