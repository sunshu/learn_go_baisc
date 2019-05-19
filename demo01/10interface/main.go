package main

/**
接口用来表示任何类型
*/

import (
	"./mock"
	"./real"
	"fmt"
	"time"
)

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retiever

	//r = &mock.Retriever{"this is imooc","Mozilla/5.0"}

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Print("%T %v\n",r,r)

	inspect(r)
	//fmt.Println(download(r))
}

func inspect(r Retiever) {
	fmt.Print("%T %v\n", r, r)
	switch v := r.(type) {
	case *real.Retriever:
		fmt.Println("Contents", v.UserAgent)
	case *mock.Retriever:
		fmt.Println("v.UserAgent", v.UserAgent)
	}
}
