package main

import "fmt"

func main(){
	fmt.Print("Enter an amount: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println("The answer is ", output)
}