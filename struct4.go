package main

import "fmt"

type test struct {
	name string
	age  int
}

func main() {
	a := &test{
		name: "wangye",
		age:  21,
	}

	A(a)
	B(a)
	fmt.Println(a)
}

func A(person *test) {
	person.name = "zhenchao"
	person.age = 25
	fmt.Println(person)
}

func B(person *test) {
	person.name = "zhenchao"
	person.age = 22
	fmt.Println(person)
}
