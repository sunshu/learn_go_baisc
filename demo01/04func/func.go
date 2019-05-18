package main

import "fmt"

func main()  {
	var c ,d = div(5,3)
	fmt.Println(c)
	fmt.Println(d)
}


func div(a,b int) (c,d int){

	return  a/b,a%b
}
