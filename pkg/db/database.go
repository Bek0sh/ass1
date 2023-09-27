package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Bek0sh/soft-ass1/pkg/db/idb"
)

type DbConnection struct {
	Con idb.IDatabase
}

var data *sql.DB

func NewDbConn(con idb.IDatabase) *DbConnection {
	return &DbConnection{Con: con}
}

type Postgres struct {
	User     string
	Password string
	Database string
}

func (p *Postgres) Connect() {
	conn := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", p.User, p.Password, p.Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	data = db
}

func (db *DbConnection) GetDb() *sql.DB {
	return data
}
