package main

import ("fmt"; "math")

type Circle struct {
	//fields
	x, y, r float64
}

//method 
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//external function
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(2, (x2 - x1)) + math.Pow(2, (y2 - y1)))
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main(){

	//to instantiate:
//	var c Circle //will have 0 fields
//	d := new(Circle) //again, 0
//	e := Circle{x: 0, y: 0, r: 5} //values
	f := Circle{1,2,3}

	//to access

	fmt.Println(f.x, f.y, f.r)

	fmt.Println("area", circleArea(&f))

	//using circles area method
	fmt.Println("area using method", f.area())

	rect := Rectangle{2,3,7,9}

	fmt.Println("area of rectangle", rect, "is", rect.area())
}
