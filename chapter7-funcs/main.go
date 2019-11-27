package main

import "fmt"

func main(){
	xs := []float64{98,93,77,82,93}

	total := 0.0
	for _, v := range xs {
		total += v
	}

	fmt.Println("Total:", total)

	fmt.Println("Average:", average(xs))

	x, y := multi()

	fmt.Println("Both returned from a single func:", x, y)

	fmt.Println("addition", add(1,2,3,4,12,3))

	//closure (inner functions)
	l := 0
	increment := func() int {
		l++ //has access to above variables
		return l
	}

	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("Factorial of 5:", factorial(5))

	deferFunc()

}


//Panic with defer
func panicDeferFunc(){
	defer func(){
		str := recover()
		fmt.Println("recovered:",str)
	}()
	panic("PANIC")
}

//panic, recover
func panicFunc(){
	//panic("PANIC")
	//this will never be hit as panic immediately stops execution.
	//str := recover()
	//fmt.Println(str)
}

//defer
func first(){
	fmt.Println("1st")
}

func second(){
	fmt.Println("2nd")
}

func deferFunc(){
	defer second()
	first()
}

//recursion
func factorial(x uint ) uint {
	if x == 0 { return 1 }

	return x * factorial(x-1)
}

//closure - returning a function.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}



//variadic (var args)
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//multiple return values!!!!
func multi() (int, int) {
	return 5, 6
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}