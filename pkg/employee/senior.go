package employee

import "log"

type Senior struct {
	Helper IEmp
}

func (s *Senior) Salary() int {
	return s.Helper.Salary() + 500000
}

func (s *Senior) Expirience() int {
	return s.Helper.Expirience() + 24
}

func (s *Senior) Info() {
	log.Printf("It is Senior dev with exp: %d Month, and salary: %d", s.Expirience(), s.Salary())
}
