package main

import (
	".."
	"fmt"
)

func main() {
	q := queue.Queue{
		1,
	}
	q.Push(123)
	q.Pop()
	fmt.Println(q.Pop())

}
