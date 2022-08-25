package main

import "fmt"

func main() {
	//带缓冲的channel 3是容量
	c := make(chan int, 3)
	fmt.Println(len(c), cap(c))

}
