package main

import "fmt"

type Test struct {
	a string
}

func (t *Test) show() {
	fmt.Println(t.a)
}

// 用來建構 Test 的假建構子
func newTest() (test *Test) {
	test = &Test{a: "foobar"}
	// 這裡會回傳一個型態是 *Test 建構體的 test 變數
	return
}

func main() {
	b := newTest()
	b.show() // 輸出：foobar
}
