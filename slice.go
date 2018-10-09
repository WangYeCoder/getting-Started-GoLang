package main

import "fmt"

func main() {
	//var ar = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '8', '0'}

	array := []int{}
	array[0] = 11
	array[1] = 33
	array[2] = 44

	sslic := []int{1, 2, 3}
	// 声明两个含有byte的slice
	//var a []byte
	//var b []byte
	//var c []byte
	//// a指向数组的第3个元素开始，并到第五个元素结束，
	//a = ar[2:5]
	//b = ar[1:7]
	//c = ar[:]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]

	// b的元素是：ar[3]和ar[4]
	fmt.Print(sslic)
}
