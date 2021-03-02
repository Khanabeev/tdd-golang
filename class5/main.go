package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	radios float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radios * c.radios
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
