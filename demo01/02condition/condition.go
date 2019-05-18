package main

import "fmt"
// import "io/ioutil"

func main() {
	const fileName = "basic.go"
	// contant,error  :=  ioutil.ReadFile(fileName)

	// if contant,error  :=  ioutil.ReadFile(fileName); error != nil {
	// 	fmt.Println(error)
	// }else{
	// 	fmt.Println("%s\n",contant)
	// }
	garde(-2);

}

func garde(score int) string{
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("错误"))	
	case score < 60:
		g = "F"
	case score <70:
		g = "D"
	case score < 80:
		g = "c"	
	case score < 90:
		g = "b"	
	case score < 95:
		g = "a"	
	}
	return g

}





