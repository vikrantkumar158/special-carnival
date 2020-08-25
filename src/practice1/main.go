//Example of struct and function overloading

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	breadth float64
}

type Cuboid struct {
	Rectangle
	height float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (cu Cuboid) Volume() float64 {
	return cu.length * cu.breadth * cu.height
}

func main() {
	var (
		r1 = Rectangle{20,10}
		c1 = Circle{7}
		cu1 = Cuboid{Rectangle{20, 10}, 10}
	)
	fmt.Println("Area of circle is:", c1.Area())
	fmt.Println("Area of rectangle is:", r1.Area())
	fmt.Println("Volume of cuboid is:", cu1.Volume())
}