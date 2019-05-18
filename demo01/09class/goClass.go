package main

import "fmt"

func  main()  {


	// var zhangsan = user{1,"zhangsan"}
	// lisi := user{2,"lisi"}
	// fmt.Println(zhangsan)
	// fmt.Println(lisi)
	n2left := node{3,nil,nil}
	n2reft := node{4,nil,nil}

	n1left := node{2,&n2left,&n2reft}
	n1 := node{1,&n1left,nil}
	fmt.Println(n1)
	n1.traverse()


}

type node struct {
	value int
	leftNode *node
	rightNode *node
}

func (root *node) traverse(){
	if root == nil {
		return
	}
	root.leftNode.traverse()
	fmt.Println(root.value)
	root.rightNode.traverse()
}

func createNode(v int) *node {
	return &node{value:v}
}

func (n n) printNode


type user struct{
	id int
	name string
}

