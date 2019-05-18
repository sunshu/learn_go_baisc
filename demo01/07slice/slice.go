package main

import "fmt"

/*
 *  切片
 *  1. 切片基本操作
 *	2. 切片添加元素
 */
func main() {
	// init1()
	// arrayappend()
	makeSlice()
}

func init1() {
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	n1 := nums[1:5]
	fmt.Println(n1)

	n2 := nums[:5]
	fmt.Println(n2)

	n3 := nums[0:]
	fmt.Println(n3)

	n1[1] = 100
	fmt.Println(n1)

	capN2 := n2[4:6]
	fmt.Println(capN2)
}

func arrayappend() {
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	n1 := nums[2:4]

	n2 := append(n1, 10)

	fmt.Println(n1)
	fmt.Println(n2)

}

func makeSlice(){
	s1 := make( []int ,16,32)

	for index := 0; index < 32; index++ {
		s1 = append(s1,index)
	}

	fmt.Println(s1)
	// fmt.Println(len(s1),cap(s1))
	// s2 := make ([]int,5)

	// copy(s1,s2)
	// fmt.Println(s2)
}
