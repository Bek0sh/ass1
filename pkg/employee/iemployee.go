package employee

type IEmp interface {
	Expirience() int
	Salary() int
	Info()
}

type EmpDecorater struct {
	IE IEmp
}

func (empB *EmpDecorater) DisplayInfo() {
	empB.IE.Info()
}
