package main

import (
	"fmt"
	"math"
)
// works without this so, this shows value receiver
// with 2 different inplementations...2x structs...

// are interfaces worth the trouble?
// type Shape interface {
// 	area() float64
// }

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// inside out inheritance
// func getArea(s Shape) float64 {
// 	return s.area()
// }

func(c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func(r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	fmt.Println("interfaces")
	circle := Circle{0,0,5}
	rect := Rectangle{2.11,2.13}
	fmt.Println(circle.area())
	fmt.Println(rect.area())
	// fmt.Println(getArea(circle))
	// fmt.Println(getArea(rect))
}
