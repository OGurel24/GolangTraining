package main

import (
	"fmt"
)

type rectangular struct {
	length1 float64
	length2 float64
}

type circle struct {
	radius float64
}

func (a rectangular) area() float64 {
	return a.length1 * a.length2
}

func (a circle) area() float64 {
	return 3.14 * a.radius
}

type shape interface {
	area() float64
}

func info(a shape) {
	fmt.Println(a.area())
}

func main() {
	myrectangular := rectangular{4, 7}
	mycircle := circle{10}
	info(mycircle)
	info(myrectangular)

}
