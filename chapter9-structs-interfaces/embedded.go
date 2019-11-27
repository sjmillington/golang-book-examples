package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk(){
	fmt.Println("Hi, my name is", p.Name)
}

//For Android IS a person, rather than have a person:

type Android struct {
	Person
	Model string
}


func main(){
	p := Person{"Steve"}

	p.Talk()

	a := Android{Person{"Mike"}, "XC-30"}

	a.Person.Talk()
}