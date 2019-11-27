package main

import "fmt"

func main(){
	//numbers
	fmt.Println("1 + 1 = ", 1+1)

	//strings
	fmt.Println("Hello World"[1]) //101 not e. 
	fmt.Println(len("Hello World"))

	//bools - all behave as expected
	fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
}