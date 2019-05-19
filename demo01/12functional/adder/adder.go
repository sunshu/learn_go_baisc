package main

import "fmt"

// 定义adder函数,他的返回是一个函数，返回的这个函数 func(int) int
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	fmt.Print("")
	a := adder()
	//for i:=0;i<10 ;i++  {
	//	fmt.Println(a(i))
	//}
	//a(5)
	a(1)
	a(2)
	a(3)
	a(4)
	a(5)
	fmt.Println(a(6))
}
