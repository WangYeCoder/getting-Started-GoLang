package main

import (
	"fmt"
)

//匿名struct
func main() {
	a := struct {
		name string
		age  int
		sex  bool
	}{
		name: "123",
		age:  1213,
		sex:  true,
	}

	fmt.Println(a)
}
