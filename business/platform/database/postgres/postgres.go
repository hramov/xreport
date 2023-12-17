package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func New(host string, port int, user, password, dbname string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return nil, errors.Wrap(err, "open postgres connection")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping postgres")
	}
	return db, nil
}
