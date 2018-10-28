package main

import "fmt"

func main() {
	var a  = 20  //声明变量

	var ip *int  //声明空指针

	ip = &a  //空指针取a的地址

	fmt.Printf("a地址是: %x\n", &a)


	fmt.Printf("指针取出来的值是: %x\n", ip)



}
