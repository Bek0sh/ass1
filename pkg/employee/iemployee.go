package employee

type IEmp interface {
	Expirience() int
	Salary() int
	Info()
}

type EmpBehabior struct {
	IE IEmp
}

func (empB *EmpBehabior) DisplayInfo() {
	empB.IE.Info()
}
