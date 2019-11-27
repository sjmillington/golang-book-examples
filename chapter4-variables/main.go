package main

import "fmt"

func main(){

	
	var x string = "Hello World"
	fmt.Println(x)

	//string equality
	var one string = "hello"
	var two string = "hello"

	fmt.Println(one == two) //true

	//short hand
	y := "Hello another world" //type inference here
	fmt.Println(y)

	f()

	//constants
	const c = "Constant World"

	//multiple

	var (
		a = 5
		b = 10
		c = 15
	)


}

//variable scope

var inScope string = "outside function scope"

func f(){
	fmt.Println(inScope)
	//fmt.Println(x) wont work
}