package employee

import "log"

type Middle struct {
	Helper IEmp
}

func (m *Middle) Expirience() int {
	return m.Helper.Expirience() + 18
}

func (m *Middle) Salary() int {
	return m.Helper.Salary() + 600000
}

func (m *Middle) Info() {
	log.Printf("It is Middle dev with exp: %d Month, and salary: %d", m.Expirience(), m.Salary())
}
