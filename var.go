package main

import "fmt"
import "io/ioutil"

func variable() {
	var a int
	var s string
	fmt.Println(a, s)
}

func varfuchuzhi() {
	var a int = 3
	var s string = "bc"
	fmt.Println(a, s)
}

func readfile() {
	const filename = "abc.txt"
	result, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func eval(op func(int,int) int,a,b int) int {
	return op(a,b)
}



func main() {
	//readfile()
}
