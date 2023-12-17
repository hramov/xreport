package main

import (
	"context"
	v1 "github.com/hramov/xreport/app/service/v1"
	"github.com/hramov/xreport/business/platform/database/postgres"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	log := log.New(os.Stdout, "XReport: ", log.Ldate|log.Ltime)

	if err := run(ctx, log); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *log.Logger) error {
	log.Println("starting XReport")
	defer log.Println("stop XReport")

	// database
	db, err := postgres.New("localhost", 5432, "postgres", "postgres", "xreport")
	if err != nil {
		return err
	}

	defer db.Close()

	// mux
	shutdownCh := make(chan struct{}, 1)
	mux := v1.New(shutdownCh, db, log)

	// server
	server := &http.Server{
		Addr:              "localhost:3000",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	serverError := make(chan error, 1)
	go func() {
		serverError <- server.ListenAndServe()
	}()

	// shutdown
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	select {
	case err = <-serverError:
		return errors.Wrap(err, "server error")
	case <-shutdownCh:
		log.Println("shutdown error")
		err = shutdownServer(server)
		if err != nil {
			return errors.Wrap(err, "shutdown error")
		}
	case <-ctx.Done():
		log.Println("shutting down")
		err = shutdownServer(server)
		if err != nil {
			return errors.Wrap(err, "shutdown error")
		}
	}

	return nil
}

func shutdownServer(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		server.Close()
		return errors.Wrap(err, "server shutdown error")
	}
	return nil
}
