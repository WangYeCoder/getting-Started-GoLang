package main

import "fmt"

type struct1 struct {
	i1 int
	f1 float32
	str string
}



func main()  {
	ms:=struct1{}
	ms.i1=10
	ms.f1=15.1
	ms.str="我真牛逼"

	fmt.Printf("the int is : %d\n",ms.i1)
	fmt.Println(ms)
	}