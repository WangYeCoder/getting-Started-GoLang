package main

import "fmt"

func main() {
	s1 := func() int {
		a := 123
		return a
	}

	fmt.Println(s1)
}
