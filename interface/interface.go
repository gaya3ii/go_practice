package main

import (
	"fmt"
	"math"
)

// Interfaces
type Shape interface {
	Area() float64
}

type Solid interface {
	Volume() float64
}

// Structs
type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Sphere struct {
	Radius float64
}

// Triangle methods
func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Rectangle methods
func (r *Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

// Sphere methods
func (s *Sphere) Area() float64 {
	return 4 * math.Pi * s.Radius * s.Radius
}

func (s *Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.Radius, 3)
}

func main() {
	// all three satisfy Shape interface
	shapes := []Shape{
		&Triangle{Base: 10, Height: 5},
		&Rectangle{Length: 4, Breadth: 6},
		&Sphere{Radius: 3},
	}

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}

	solids := []Solid{
		&Sphere{Radius: 3},
	}
	for _, solid := range solids {
		fmt.Printf("Volume: %.2f\n", solid.Volume())
	}
}
