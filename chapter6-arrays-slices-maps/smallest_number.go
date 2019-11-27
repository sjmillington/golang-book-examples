package main

import "fmt"

func main(){

	x := []int8{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

  fmt.Println("Smallest number is", smallestNumber(x))

}

func smallestNumber(arr []int8 ) int8 {

	var y int8 = 1 << 7 - 1

	for _, value := range arr {
		if value < y {
			y = value
		}
	}

	return y
}