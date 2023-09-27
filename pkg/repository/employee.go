package repository

import (
	"database/sql"

	"github.com/Bek0sh/soft-ass1/pkg/models"
)

type IRepository interface {
	CreateUser(*models.Employee) (int, error)
	GetUserByName(name string) (*models.Employee, error)
	DeleteUserWithId(int) error
}

type repo struct {
	db *sql.DB
}
