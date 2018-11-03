package main

import "fmt"

type Rectangle struct {
	with, height float32
}

type Circle struct {
	radius float64
}

func (r *Rectangle) area() float32 {
	fmt.Println(r)
	return r.height * r.with
}

func (c *Circle) area() float64 {
	return c.radius * c.radius
}

func main() {
	r1 := &Rectangle{
		with:   10,
		height: 20,
	}

	r2 := &Rectangle{
		with:   20,
		height: 10,
	}

	c1 := &Circle{
		radius: 100,
	}

	c2 := &Circle{
		radius: 50,
	}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())

	fmt.Println("Area of r1 is: ", c1.area())
	fmt.Println("Area of r2 is: ", c2.area())
}
