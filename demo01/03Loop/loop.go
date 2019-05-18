package main

import "fmt"


func  main()  {
	fmt.Println(12/2)
	fmt.Println(13%5)
	count := 100
	result := 0
	for i := 0; i <= count; i++ {
		result = i+result
	}
	fmt.Println(result)
}


// func converToBin(n int) string{
// 	// result := ""
// 	// for  ;  n > 0;  n/=2 {
// 	// 	lsb := n%2
// 	// }



// 	return ""
// }