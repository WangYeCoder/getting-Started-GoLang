package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{
		1: "wangye",
		2: "boy",
		3: "188",
		4: "yeah",
	}

	i := 0
	s := make([]int, len(m))

	for k, _ := range m {
		s[i] = k
		i++
	}

	sort.Ints(s)
	fmt.Println(s)

	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	//
	//slciemap := make([]map[int]string, 5)
	//
	//for k,_ := range slciemap {
	//	slciemap[k] = make(map[int]string, 1)
	//	slciemap[k][1] = "ok"
	//	fmt.Println(slciemap[k])
	//}
	//fmt.Println(slciemap)

}
