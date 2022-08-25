package main

import (
	"fmt"
)

func main() {
	//定义一个channel
	ch := make(chan int)

	//创建异步goroutine
	go func() {
		defer fmt.Println("goroutine defer")
		fmt.Println("goroutine start")
		//向channel中写入数据 必须传进去,如果没有传进去则会报错,这里会阻塞,等待外层取出才会继续执行
		ch <- 1
		fmt.Println("goroutine end")
	}()

	//从channel中取出ch变量数据, 这里会阻塞,等待ch被写入数据,
	num := <-ch
	fmt.Println("接收到的数据", num)

}
