package main

import "fmt"
import "math"

func main() {
	initVar()
	consts() 
	enums()
}

// 定义变量
func initVar(){
	a,b,c := true,3,"Hello world"
	fmt.Println(a,b,c)
}

// 定义常量
func consts(){
	const(
		ac = 1
		ae = 2.0
	)
	const fileName = "20190501img.jpg"
	const a,b = 3,4
	var c  int 
	c  = int(math.Sqrt(a*a+b*b))
	fmt.Println(a,b,c)
}

// 枚举
func enums(){
	const(
		cpp = 1 <<  (iota * 10)
		php
		java 
		golang 
		javascript
	)
	fmt.Println(cpp,java,golang,javascript)
}



