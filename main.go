package main

import (
	"database/sql"

	"github.com/Bek0sh/soft-ass1/pkg/config"
	"github.com/Bek0sh/soft-ass1/pkg/db"
	"github.com/Bek0sh/soft-ass1/pkg/employee"
)

var pDB *sql.DB
var mapa map[int]employee.Employee

func init() {
	cfg := config.GetConfig()

	postgres := &db.Postgres{
		User:     cfg.Postgre.DbUsername,
		Password: cfg.Postgre.DbPassword,
		Database: cfg.Postgre.DbName,
	}

	cache := &db.Cache{}

	data := db.NewDbConn(postgres)
	cacheData := db.NewDbConn(cache)

	data.Con.Connect()
	cacheData.Con.Connect()

	pDB = data.GetPostgre()
	mapa = cacheData.GetCache()
}

func main() {

	jun := &employee.Junior{
		Helper: &employee.Employee{},
	}
	mid := &employee.Middle{
		Helper: &employee.Junior{
			Helper: &employee.Employee{},
		},
	}
	senior := &employee.Senior{
		Helper: &employee.Middle{
			Helper: &employee.Junior{
				Helper: &employee.Employee{},
			},
		},
	}
	dev1 := employee.EmpBehabior{IE: jun}
	dev2 := employee.EmpBehabior{IE: mid}
	dev3 := employee.EmpBehabior{IE: senior}

	dev1.DisplayInfo()
	dev2.DisplayInfo()
	dev3.DisplayInfo()

}
