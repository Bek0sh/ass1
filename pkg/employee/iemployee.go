package employee

type IEmp interface {
	Expirience() int
	Salary() int
	Info()
}

type EmpDecorater struct {
	IE IEmp
}

func (empB *EmpBehabior) DisplayInfo() {
	empB.IE.Info()
}
