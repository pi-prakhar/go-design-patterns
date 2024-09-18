package factory

//getting Employee from other file.
/**
type Employee struct {
	Name, Position string
	AnnualIncome   int
}
**/

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}

}
