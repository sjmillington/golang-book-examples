package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string){
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string){
	for i := 0; ; i++{
		c <- "pong"
	}
}

func printer(c <-chan string){
	for {
		fmt.Println(<- c)
		time.Sleep(time.Second)
	}
}

func main(){
	var c chan string = make(chan string)


	go pinger(c)
	go printer(c)
	go ponger(c)
	
	var input string
	fmt.Scanln(&input)
}