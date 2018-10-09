package main

import "fmt"

var quanju = 123

func main() {
	s := `hel

lo`
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
}
