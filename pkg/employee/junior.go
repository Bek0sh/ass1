package employee

import "log"

type Junior struct {
	Helper IEmp
}

func (j *Junior) Expirience() int {
	return j.Helper.Expirience() + 6
}

func (j *Junior) Salary() int {
	return j.Helper.Salary() + 200000
}

func (j *Junior) Info() {
	log.Printf("It is Junior dev with exp: %d Month, and salary: %d", j.Expirience(), j.Salary())
}
