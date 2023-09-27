package main

import (
	"log"

	"github.com/Bek0sh/soft-ass1/pkg/db"
	"github.com/Bek0sh/soft-ass1/pkg/models"
)

func init() {
	cache := &db.Cashe{}

	data := db.NewDbConn(cache)
	data.Con.Connect()
	database := data.GetCashe()
	database.M[1] = models.Employee{Name: "Beka"}
	log.Print(database.M[1].Name)
}

func main() {

}
