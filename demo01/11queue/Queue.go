package main

import "fmt"

func main() {
	q := Queue{1, "2", true}
	q.Push("abv")
	q.Pop()
	fmt.Println(q)
	fmt.Print(q.isEmpty())
}

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
