package main

import "fmt"

func main(){
	forOneWay()
	forInOne()
	ifDemo()
	switches()
}

//go has no while, do, until or forEach... only for...

func forOneWay(){
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func forInOne(){
	for i := 1; i <= 10 ;i++ {
		fmt.Println(i)
	}
}

func ifDemo(){
	for i := 1; i <= 10 ;i++ {
		if i % 2 == 0 {
			fmt.Println("even")
		}else{
			fmt.Println("odd")
		}
	}
}

func switches(){
	for i := 1; i <= 10 ;i++ {
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		default: fmt.Println("Other")
		}
	}
}