package main

import "fmt"

//创建斐波那契数列的函数
func fibonacci(ch1, ch2 chan int) {
	x, y := 0, 1
	for {
		select {
		//检测chan2是否可读
		case <-ch2:
			fmt.Println("read from ch2")
			return
		//检测chan1是否可写
		case ch1 <- x:
			x, y = y, x+y
			fmt.Println("write ch1")
		}
	}
}

func main() {

	//创建一个channel
	ch := make(chan int)
	//创建一个channel
	ch2 := make(chan int)

	//子协程
	go func() {
		//读取channel1数据
		for i := 0; i < 10; i++ {
			data := <-ch
			fmt.Println("读取 data=", data)
		}

		//写入channel2数据
		ch2 <- 1
		fmt.Println("写入 ch2")
	}()

	//写个函数监听两个channel的状态
	fibonacci(ch, ch2)

}
