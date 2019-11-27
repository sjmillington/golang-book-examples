package main

import (
	"fmt"
	"hash/crc32"
	"crypto/sha1"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    return 0, err
  }
  h := crc32.NewIEEE()
  h.Write(bs)
  return h.Sum32(), nil
}

func main(){
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("./test1.txt")
  if err != nil {
		fmt.Println(err)
    return
  }
  h2, err := getHash("./test2.txt")
  if err != nil {
		fmt.Println(err)
    return
  }
	fmt.Println(h1, h2, h1 == h2)
	
	fmt.Println("SHA1 password hashing")
	sha := sha1.New()
	sha.Write([]byte("test"))
	bs := sha.Sum([]byte{})
	fmt.Println(bs)
}

