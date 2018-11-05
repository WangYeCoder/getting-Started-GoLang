package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h *Human) String() string {
	return "❰" + (*h).name + " - " + strconv.Itoa((*h).age) + " years -  ✆ " + (*h).phone + "❱"
}

func main() {

	hun := &Human{
		"!@3",
		123,
		"123",
	}

	fmt.Println(hun.String())

}
