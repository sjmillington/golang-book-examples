package main

import "fmt"

func main(){
	fmt.Print("Enter a fahrenheit to convert: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9

	fmt.Println(input, " fahrenheit converts to ", output, " celcius")
}