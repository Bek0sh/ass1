package main

import (
	"database/sql"

	"github.com/Bek0sh/soft-ass1/pkg/db"
	"github.com/Bek0sh/soft-ass1/pkg/employee"
)

var pDB *sql.DB

func init() {
	postgres := &db.Postgres{
		User:     "postgres",
		Password: "1234",
		Database: "todoauth",
	}

	data := db.NewDbConn(postgres)
	data.Con.Connect()
	pDB = data.GetPostgre()
}

func main() {
	emp := &employee.Employee{}
	jun := &employee.Junior{Helper: emp}
	mid := &employee.Middle{Helper: &employee.Junior{Helper: emp}}

	dev := employee.EmpBehabior{IE: emp}
	dev1 := employee.EmpBehabior{IE: jun}
	dev2 := employee.EmpBehabior{IE: mid}

	dev.DisplayInfo()
	dev1.DisplayInfo()
	dev2.DisplayInfo()

}
