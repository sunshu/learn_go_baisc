package main

import "fmt"

func main()  {
	var childs [5]int
	printArray(childs);
	girls := [5]int{1,2,3,4}
	printArray(girls);
	boys := []int{1,2,3,4}
	printItem(boys);

	man := []int{222,223,224,225,226,227}
	printItemPoint(&man)
}


func printArray( childs [5]int ){
	for i,v := range childs{
		fmt.Println(i,v);
	}
}

func printItem( childs []int ){
	for i,v := range childs{
		fmt.Println(i,v);
	}
}

func printItemPoint( childs *[]int ){
	for i,v := range *childs{
		fmt.Println(i,v);
	}
}