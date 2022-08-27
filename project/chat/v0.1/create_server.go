package main

import (
	"fmt"
	"net"
)

//定义一个服务结构体
type Server struct {
	Ip   string
	Port int
}

//启动服务
// func (s *Server) 这种是类方法的写法
func (s *Server) Start() {
	fmt.Println("Server start")
	//监听ip和port的TCP链接赋值给listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	//关闭listener
	defer listener.Close()

	//循环监听
	for {
		//等待客户端链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		//处理客户端链接, 异步处理
		go s.handleConn(conn)
	}

}

//异步处理客户端链接, 形参conn是一个net.Conn类型的变量
//这是类方法
func (s *Server) handleConn(conn net.Conn) {
	fmt.Println("handleConn")
}

// NewServer 创建一个服务对外提供的方法
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}
