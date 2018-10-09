package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("输出一个错误吧")
	if err != nil{
		fmt.Print("并没有报错")
	}
}
