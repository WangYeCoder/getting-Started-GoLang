package main

import "fmt"

func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func main() {
	x := 3
	y := 4
	thana, thanb := SumAndProduct(x, y)
	fmt.Print(thana, thanb)
}
