package main

import "fmt"

func main(){
	arrays()
	slices()
	maps()

	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(x[2:5])

}

func maps(){

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println("map: ", x)
	fmt.Println("key: ", x["key"])

	delete(x, "key")

	fmt.Println("map after delete: ", x)

	elements := make(map[string]string)
	elements["H"] = "Hydrogren"
	elements["He"] = "Helium"
	elements["Au"] = "Gold"

	name, ok := elements["He"]

	fmt.Println(name, ok)

	unName, unOk := elements["Un"]

	fmt.Println(unName, unOk)

	//we can use the success value for conditionals
	if auName, auOk := elements["Au"]; auOk {
		fmt.Println(auName, "IS IN!")
	}


	//NESTING

	elementsObjectMap := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
	}

	if el, ok := elementsObjectMap["H"]; ok {
    fmt.Println(el["name"], el["state"])
  }
}

func slices(){
	//segment of an array
	x := make([]float64, 5, 10) //5-10 of the array.

	fmt.Println("Make function", x);
	//short hand.
	arr := [5]float64{1,2,3,4,5}
	y := arr[0:2]

	fmt.Println("short hand slice", y)

	//functions - append
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("append", slice1, slice2)


	slice3 := []int{1,2,3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("copy", slice3, slice4)
	
}

func arrays(){

		var x [5]int
		x[4] = 100
		fmt.Println("simple array", x)

		var total float64 = 0
		//we need to cast to the same data type.
		//shorthand since we don't use i.
		for _, value := range x{
			total += float64(value)
		}
		fmt.Println("Average", total / float64(len(x)))

		//short hand array creation
		y := [5]float64{ 98, 93, 92, 91.0, 88 }

		fmt.Println("Short hand creation", y)

		//also allowed on broken lines - trailing comma needed
		p := [5]int{
			2,
			3,
			4,
		}

		fmt.Println("Broken lines", p)
}