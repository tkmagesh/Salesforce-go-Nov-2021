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

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type ShapeWithArea interface {
	Area() float64
}

func main() {
	c := Circle{Radius: 12}
	//fmt.Println("Area : ", c.Area())
	PrintArea(c)
	r := Rectangle{Width: 12, Height: 10}
	//fmt.Println("Area : ", r.Area())
	PrintArea(r)
}

//utility functions
func PrintArea(sa ShapeWithArea) {
	fmt.Println("Area : ", sa.Area())
}
