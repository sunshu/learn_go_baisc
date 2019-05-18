package main

import "fmt"

func main()  {
	
	m := map[string]string{
		"id":"1",
		"name":"Licy",
		"sex":" ",
	}

	// fmt.Println(m)

	// for key,_ := range m {
	// 	fmt.Println(key)
	// }

	for key,_ := range m {
		if name,ok := m["name"]; ok{
			fmt.Println(key)
			fmt.Println(name)
		}

		// fmt.Println(key)

	}
}