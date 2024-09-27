package prototype

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

// Circle is a concrete prototype

type Circle struct {
	Radius int
	Color  string
}

func (c *Circle) Clone() Cloneable {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle [radius: %d, color: %s]", c.Radius, c.Color)
}

//Rectangle is another concrete prototype

type Rectangle struct {
	Width  int
	Height int
	Color  string
}

func (r *Rectangle) Clone() Cloneable {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}
