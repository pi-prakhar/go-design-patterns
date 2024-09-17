package factory

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{Position: position, AnnualIncome: annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{
		name,
		f.Position,
		f.AnnualIncome,
	}
}
