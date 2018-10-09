package main

import "fmt"

func main() {
	x := 1000

	if y := 1000; x > y {
		fmt.Print("x is greater than ")
	} else if x < y {
		fmt.Print("x is less than 10")
	} else {
		fmt.Print("x=y")
	}
}

//fmt.print
