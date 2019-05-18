package main

import "fmt"

func main(){
	var k int
	var i int
	for ; i < 5; i++ {
		for k = 1; k <= (2 * i - 3) ; k++{		
		fmt.Printf("*")
	}
}
}