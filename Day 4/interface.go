package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Sqaure struct {
	side float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}


func (s Sqaure) Area() float64{
	return s.side * s.side
}

func (s Sqaure) Perimeter() float64{
	return 4 * s.side
}

func PrintInfo(a Shape){
	fmt.Printf("Area of %T is %.2f while its perimeter is %.2f\n",a,a.Area(),a.Perimeter())
}
