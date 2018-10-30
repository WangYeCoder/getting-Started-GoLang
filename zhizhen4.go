package main

import "fmt"

func main() {
	i1 := 5
	var i2 *int
	i2 = &i1
	fmt.Println("An integer: %d, it's location in memory: %p\n", i2, i1, &i1)
}
