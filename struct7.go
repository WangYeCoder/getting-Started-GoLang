package main

import "fmt"

type human struct {
	name   string
	age    int
	weight int
}

type student struct {
	human
	name   string
	specil string
}

func (name *student) changname() {
	name.name = "huohuohuo"
	fmt.Println(name)
}

func main() {
	mark := student{
		human{
			name:   "123",
			age:    123,
			weight: 123,
		},
		"6666",
		"123123",
	}

	fmt.Println(mark.name)
	mark.changname()
	fmt.Println(mark.name)

}
