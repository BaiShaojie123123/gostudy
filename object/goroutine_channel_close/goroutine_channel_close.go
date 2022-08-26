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

	//第一种 主协程死循环读取数据
	//如果channel没有关闭,那么for循环会一直阻塞,直到channel关闭,才会退出for循环,如果没写close(ch)则会一直阻塞报错
	for {
		//从channel中读取数据
		// data, ok := <-ch; 表示从channel中读取数据,如果channel没有关闭,ok为true,如果channel已经关闭,ok为false
		// if data, ok := <-ch; ok == true { 表示先执行data, ok := <-ch;得到data和ok的值,然后判断ok的值,如果ok为true,则执行if语句

		//if data, ok := <-ch; ok == true {
		//与上面的写法等价
		if data, ok := <-ch; ok {
			fmt.Println("data=", data)
		} else {
			break
		}
	}

	//第二种 主协程死循环读取数据
	//range ch 表示尝试从channel中读取数据,range阻塞等待channel中有数据,如果channel没有关闭,那么for循环会一直阻塞,直到channel关闭,才会退出for循环,如果没写close(ch)则会一直阻塞报错
	//for data := range ch {
	//	fmt.Println("data=", data)
	//}

}
