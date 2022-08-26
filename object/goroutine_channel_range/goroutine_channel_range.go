package main

import "fmt"

func main() {
	//创建一个无缓冲channel
	ch := make(chan int)
	//子协程
	go func() {
		//循环写入数据
		for i := 0; i < 10; i++ {
			ch <- i
		}

		//关闭channel,关闭后,channel不能再写入数据,但是可以继续读取数据
		close(ch)

	}()

	//主协程死循环读取数据
	//range ch 表示尝试从channel中读取数据,range阻塞等待channel中有数据,如果channel没有关闭,那么for循环会一直阻塞,直到channel关闭,才会退出for循环,如果没写close(ch)则会一直阻塞报错
	for data := range ch {
		fmt.Println("data=", data)
	}

}
