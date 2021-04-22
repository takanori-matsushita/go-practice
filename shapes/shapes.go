package main

import (
	"math"
)

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

// func Area(rectangle Rectangle) float64 {
// 	return
// }
