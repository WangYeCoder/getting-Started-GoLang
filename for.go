package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i > 0; i++ {
		sum += i
		if sum > 100000000000 {
			break
		}
		fmt.Print(sum, "\n")
	}
}
