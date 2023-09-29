package employee

import (
	"log"
)

var Number int

type Employee struct {
}

func (e *Employee) Expirience() int {
	return 1
}

func (e *Employee) Salary() int {
	return 100000
}

func (e *Employee) Info() {
	log.Printf("It is simple Employee with exp: %d Month, and salary: %d", e.Expirience(), e.Salary())
}
