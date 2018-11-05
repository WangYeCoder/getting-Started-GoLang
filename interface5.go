package main

import "fmt"

type Retiver interface {
	Get(url string) string
}

func Download(r Retiver) string {
	return r.Get("wwww.wangyewoaini.com")
}

func main() {
	var r Retiver
	fmt.Println(Download(r))
}
