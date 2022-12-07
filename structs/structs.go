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
	Radius  float64
}

type Triangle struct {
	Base float64
	Height float64
}

func Perimeter(rec Rectangle) float64 {
	result := (rec.Width + rec.Height) * 2
	return result
}

func (r Rectangle) Area() float64 {
	rectangleArea := (r.Width * r.Height)
	return rectangleArea
}

func (c Circle) Area() float64 {
	circleArea := (math.Pi * c.Radius * c.Radius)
	return circleArea
}

func (t Triangle) Area() float64 {
	triangleArea := (t.Base * t.Height) / 2
	return triangleArea
}
