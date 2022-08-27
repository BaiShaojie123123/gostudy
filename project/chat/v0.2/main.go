package main

func main() {
	//创建服务
	server := NewServer("127.0.0.1", 8080)
	//启动服务
	server.Start()
}
