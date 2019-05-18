package main

import "fmt"

func main()  {
	
	var a int = 2
	var pa *int = &a
		 
	*pa = 3
	fmt.Println(a)

	i := 6
	j := 7
	// fmt.Println(&i,&j)
	exchangeInt(&i,&j)
	// fmt.Println(&i,&j)
	fmt.Println(i,j)

	i,j =  switchab(i,j)
	fmt.Println(i,j)


}

func exchangeInt(a,b *int) {
	// fmt.Println(*a,*b)
	*b,*a  = *a,*b
	// fmt.Println(a,b)
	// fmt.Println(*a,*b)
}

func switchab(a ,b int)(int, int) {
	return b,a 
}