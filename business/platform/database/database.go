package database

import (
	"fmt"

	"github.com/hramov/xreport/business/platform/database/postgres"
	"github.com/jmoiron/sqlx"
)

type Db string

const (
	Postgres Db = "pg"
	MSSql    Db = "mssql"
)

func New(driver Db, host string, port int, user, password, dbname string) (*sqlx.DB, error) {
	switch driver {
	case Postgres:
		return postgres.New(host, port, user, password, dbname)
	case MSSql:
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf("unknown driver")
	}
}
