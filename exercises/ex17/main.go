package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (circle Circle) Area() float64 {
	return math.Pow(circle.radius, 2) * math.Pi
}

func main() {
	var rect = Rectangle{20, 10}
	var circle = Circle{10}

	fmt.Printf("The area of the rectangle where width is %v and height is %v is %v.\n", rect.width, rect.height, rect.Area())
	fmt.Printf("The area of the circle where radius is %v is %v.\n", circle.radius, circle.Area())
}
