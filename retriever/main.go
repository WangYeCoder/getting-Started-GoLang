package main

import (
	"./baidu"
	"./real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func Download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case baidu.Retriever:
		fmt.Println("content:", v.Contents)
	case *real.Retriever:
		fmt.Println("useragent:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = baidu.Retriever{
		"this is baidu.com",
	}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)
	//fmt.Println(Download(r))
}
