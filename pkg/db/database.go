package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/Bek0sh/soft-ass1/pkg/db/idb"
	"github.com/Bek0sh/soft-ass1/pkg/employee"
)

type DbConnection struct {
	Con idb.IDatabase
}

var data *sql.DB

var mapa map[int]employee.Employee

func NewDbConn(con idb.IDatabase) *DbConnection {
	return &DbConnection{Con: con}
}

type Postgres struct {
	User     string
	Password string
	Database string
}

type Cache struct {
}

func (p *Postgres) Connect() {
	conn := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", p.User, p.Password, p.Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	data = db
	log.Print("conn success")
}

func (c *Cache) Connect() {
	mapa = make(map[int]employee.Employee)
}

func (db *DbConnection) GetPostgre() *sql.DB {
	return data
}

func (db *DbConnection) GetCache() map[int]employee.Employee {
	return mapa
}
