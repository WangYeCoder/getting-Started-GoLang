package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}

func copySlice(somearr []int, itarr []int) {
	copy(itarr, somearr)
	fmt.Printf("len=%d,cap=%d\n", len(itarr), cap(itarr))

}

func deleSlice(one []int, two []int) {

}

func main() {
	//定义一个slice
	//var s [] int

	//for i := 0; i < 400; i++ {
	//printSlice(s)
	//s = append(s, i*2+1)
	//}

	s1 := []int{2, 3, 45, 6}
	//printSlice(s1)
	//s2 := []int{}

	s3 :=s1[len(s1)+1:]

	fmt.Println(s3)

}
