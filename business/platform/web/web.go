package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"syscall"
)

type App struct {
	mux      *httptreemux.ContextMux
	gm       []Middleware
	shutdown chan os.Signal
	db       *sqlx.DB
	log      *log.Logger
}

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error)

type Middleware func(Handler) Handler

func New(shutdownCh chan os.Signal, db *sqlx.DB, log *log.Logger) *App {
	return &App{
		mux:      httptreemux.NewContextMux(),
		shutdown: shutdownCh,
		db:       db,
		log:      log,
	}
}

func (a *App) Handle(method string, path string, handler Handler, lm ...Middleware) {
	a.mux.Handle(method, path, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		handler = wrapMiddleware(lm, handler)
		handler = wrapMiddleware(a.gm, handler)

		data, err := handler(ctx, w, r)
		if err != nil {
			var appErr AppError
			if errors.As(err, &appErr) {
				err = a.SendError(w, appErr.Unwrap(), appErr.Code)
			} else {
				if validateError(err) {
					a.Shutdown()
					return
				}
				err = a.SendError(w, fmt.Errorf("internal server error"), http.StatusInternalServerError)
			}

			if err != nil {
				a.shutdown <- syscall.SIGTERM
			}
			return
		}

		if data != nil {
			err = a.SendResponse(w, data, http.StatusOK)
			if err != nil {
				a.shutdown <- syscall.SIGTERM
			}
		} else {
			w.WriteHeader(http.StatusOK)
		}
	})
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// EnableCORS enables CORS preflight requests to work in the middleware. It
// prevents the MethodNotAllowedHandler from being called. This must be enabled
// for the CORS middleware to work.
func (a *App) EnableCORS(mw Middleware) {
	a.gm = append(a.gm, mw)

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		return a.SendResponse(w, "OK", http.StatusOK)
	}

	a.mux.OptionsHandler = func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		_ = handler(r.Context(), w, r)
	}
}

func (a *App) Use(middleware ...Middleware) {
	a.gm = append(a.gm, middleware...)
}

func (a *App) Shutdown() {
	a.shutdown <- syscall.SIGTERM
}

// wrapMiddleware creates a new handler by wrapping middleware around a final
// handler. The middlewares' Handlers will be executed by requests in the order
// they are provided.
func wrapMiddleware(mw []Middleware, handler Handler) Handler {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		mwFunc := mw[i]
		if mwFunc != nil {
			handler = mwFunc(handler)
		}
	}

	return handler
}

// validateError validates the error for special conditions that do not
// warrant an actual shutdown by the system.
func validateError(err error) bool {

	// Ignore syscall.EPIPE and syscall.ECONNRESET errors which occurs
	// when a write operation happens on the http.ResponseWriter that
	// has simultaneously been disconnected by the client (TCP
	// connections is broken). For instance, when large amounts of
	// data is being written or streamed to the client.
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://gosamples.dev/broken-pipe/
	// https://gosamples.dev/connection-reset-by-peer/

	switch {
	case errors.Is(err, syscall.EPIPE):

		// Usually, you get the broken pipe error when you write to the connection after the
		// RST (TCP RST Flag) is sent.
		// The broken pipe is a TCP/IP error occurring when you write to a stream where the
		// other end (the peer) has closed the underlying connection. The first write to the
		// closed connection causes the peer to reply with an RST packet indicating that the
		// connection should be terminated immediately. The second write to the socket that
		// has already received the RST causes the broken pipe error.
		return false

	case errors.Is(err, syscall.ECONNRESET):

		// Usually, you get connection reset by peer error when you read from the
		// connection after the RST (TCP RST Flag) is sent.
		// The connection reset by peer is a TCP/IP error that occurs when the other end (peer)
		// has unexpectedly closed the connection. It happens when you send a packet from your
		// end, but the other end crashes and forcibly closes the connection with the RST
		// packet instead of the TCP FIN, which is used to close a connection under normal
		// circumstances.
		return false
	}

	return true
}
