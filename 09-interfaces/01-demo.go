package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type ShapeWithArea interface {
	Area() float64
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func main() {
	c := Circle{Radius: 12}
	//fmt.Println("Area : ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)
	r := Rectangle{Width: 12, Height: 10}
	//fmt.Println("Area : ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}

//utility functions
func PrintArea(sa ShapeWithArea) {
	fmt.Println("Area : ", sa.Area())
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Println("Perimeter : ", sp.Perimeter())
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}
