package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(){
	//listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//handle the connection asyncronously
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client(){
	//connect
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send a message
	msg := "Hello world!"
	fmt.Println("sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main(){
	go server()
	go client()

	fmt.Println("Starting server")

	var input string
	fmt.Scanln(&input)
}