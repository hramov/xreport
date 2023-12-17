package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func New(host string, port int, user, password, dbname string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return nil, errors.Wrap(err, "open postgres connection")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping postgres")
	}
	return db, nil
}
