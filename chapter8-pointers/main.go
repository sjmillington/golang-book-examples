package main

import "fmt"

//* and & operators

//pass by value
func zero(x int ){
	x = 0
}

//pass by reference
func zeroByReference(xPtr *int){ //pass in the reference to an int (*int)
	*xPtr = 0 //de-reference the pointer by using * again.
}

func one(yPtr *int){
	*yPtr = 1
}

func swap(xPtr *int, yPtr *int){
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}

func main(){
	x := 5
	zero(x)
	fmt.Println(x) //still 5

	zeroByReference(&x) // find the address of x, using &
	fmt.Println(x) //now 0

	//getting a pointer using the new function
	yPtr := new(int) //value will be 0 initiallty
	one(yPtr)
	fmt.Println(*yPtr) //need to de-reference here too

	//swap
	a := 1
	b := 2

	fmt.Println("a before", a, "b before", b)
	swap(&a, &b)
	fmt.Println("a now", a, "b now", b)
}