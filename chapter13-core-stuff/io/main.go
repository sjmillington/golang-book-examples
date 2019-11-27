package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main(){
	//The long way:
	// file, err := os.Open("./test.txt")
	// if err != nil {
	// 	return
	// }

	// defer file.Close()

	// //to get the size
	// stat, err := file.Stat()
	// if err != nil {
	// 	return
	// }

	// //read it in
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil{
	// 	return
	// }

	// str := string(bs)
	// fmt.Println(str)

	//create the file
	file, err := os.Create("./test.txt")
	if err != nil{
		return
	}
	defer file.Close()

	file.WriteString("testing!")


	//read the file
	bs, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	//directory contents
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil{
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	//recursive walking
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}