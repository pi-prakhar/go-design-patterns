package builder

type BreadType string
type PattyType string

var (
	MultiGrain BreadType = "multigrain"
	WholeWheat BreadType = "whole-wheat"
)

var (
	AlooPatty    PattyType = "aloo-patty"
	ChickenPatty PattyType = "chicken-patty"
)

type Burger struct {
	BreadType BreadType
	PattyType PattyType
	Cheese    bool
	Mayonaise bool
}

type BurgerBuilder struct {
	burger *Burger
}

func NewBurgerBuilder() *BurgerBuilder {
	burger := Burger{}
	return &BurgerBuilder{&burger}
}

func (b *BurgerBuilder) ChooseBread(bread BreadType) *BurgerBuilder {
	b.burger.BreadType = bread
	return b
}

func (b *BurgerBuilder) ChoosePatty(patty PattyType) *BurgerBuilder {
	b.burger.PattyType = patty
	return b
}

func (b *BurgerBuilder) AddCheese() *BurgerBuilder {
	b.burger.Cheese = true
	return b
}

func (b *BurgerBuilder) AddMayonaise() *BurgerBuilder {
	b.burger.Mayonaise = true
	return b
}

func (b *BurgerBuilder)  Build() *Burger {
	return b.burger
}
