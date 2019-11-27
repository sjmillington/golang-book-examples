package main

import ("fmt"; "container/list"; "sort")

type Person struct {
	Name string
	Age int
}

//needs to implement the sort.Interface by implementing Len, Less and Swap
type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int){
	this[i], this[j] = this[j], this[i]
}

func main(){
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e=e.Next(){
		fmt.Println(e.Value.(int))
	}

	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)
}