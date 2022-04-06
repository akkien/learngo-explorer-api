package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sqlx.DB
}

func (pg *Postgres) Connect(connStr string) {
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	pg.DB = db
}
