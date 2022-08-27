package main

import (
	"fmt"
	"net"
	"sync"
)

//定义一个服务结构体
type Server struct {
	Ip   string
	Port int

	//用户在线列表
	//map[用户名]用户对象
	UserConns map[string]*User
	//锁
	UserConnsLock sync.RWMutex

	//消息广播队列
	MsgChan chan string
}

//新增监听一个消息广播队列方法的子协程
func (s *Server) ListenMsg() {
	//循环监听
	for {
		//将消息发送给所有在线用户
		msg := <-s.MsgChan

		//将消息发送给所有在线用户的消息队列
		s.UserConnsLock.RLock()
		for _, conn := range s.UserConns {
			conn.Msg <- msg
		}
		s.UserConnsLock.RUnlock()
	}
}

//启动服务 类方法
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

	//新增监听一个消息广播队列方法的子协程
	go s.ListenMsg()

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
	//修改

	//接收到用户链接后，创建一个用户对象
	user := NewUser(conn)
	//将用户对象添加到用户在线列表中
	s.UserConnsLock.Lock()
	s.UserConns[user.Name] = user
	s.UserConnsLock.Unlock()

	//将用户上线消息发送给所有在线用户,将消息写入广播队列
	s.MsgChan <- fmt.Sprintf("%s上线了\n", user.Name)

	//阻塞当前协程，等待用户消息
	select {}
}

// NewServer 创建一个服务对外提供的方法
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
		//初始化用户在线列表
		UserConns: make(map[string]*User),
		//初始化消息广播队列
		MsgChan: make(chan string),
	}
}
