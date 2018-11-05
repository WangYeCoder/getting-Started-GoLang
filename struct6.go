package main

import (
	f "fmt"
)

type Person struct {
	age  int
	name string
}

func Calculation(p1, p2 *Person) (value int) {
	if p1.age > p2.age {
		p1.age++
		p2.age++
		return p1.age - p2.age
	}
	p1.age++
	p2.age++
	return p2.age - p1.age
}

func main() {

	var tom Person

	tom.age = 15
	tom.name = "tom"

	bob := Person{
		age:  11,
		name: "bob",
	}

	agevalue := Calculation(&tom, &bob)

	f.Println(agevalue)

	f.Println(tom, bob)
}
