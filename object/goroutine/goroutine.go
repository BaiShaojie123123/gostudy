package main

import (
	"fmt"
	"time"
)

//创建匿名协程
func main() {
	//go 开头创建子协程 子goroutine ,与main主程序代码异步
	go func(a int) bool {
		fmt.Println("hello", a)
		return true
	}(1)
	//如果想接收子协程的参数需要用到channel

	//主程序需要等待否则上面的无法执行,因为主程序执行完毕直接就关闭了 主goroutine
	for {
		time.Sleep(1 * time.Second)
	}

}
