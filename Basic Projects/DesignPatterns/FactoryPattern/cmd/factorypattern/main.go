package main

import "fmt"

// This demostrates the factory pattern in go.
// Creates objects based on the type..

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Square struct {
	Side float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func NewShape(shape string) Shape {
	switch shape {
	case "Rectange":
		return Rectangle{Length: 10, Breadth: 20}
	case "Square":
		return Square{Side: 10}
	default:
		return nil
	}
}

func main() {
	shape := NewShape("Rectangle")
	fmt.Println(shape.Area())
}
