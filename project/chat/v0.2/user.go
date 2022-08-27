package main

import (
	"fmt"
	"net"
)

type User struct {
	//用户名
	Name string
	//用户ip
	Addr string
	//用户链接
	Conn net.Conn
	//用户消息channel
	Msg chan string
}

//用户消息处理协程,类方法
func (u *User) handleMsg() {
	//循环读取用户消息
	for {
		//读取用户消息通道
		msg := <-u.Msg
		//将用户通道中的新消息写入到用户链接中,返回给客户端
		write, err := u.Conn.Write([]byte(msg))
		if err != nil {
			return
		}
		fmt.Println("write:", write)
	}
}

//创建一个用户对外提供的方法
func NewUser(conn net.Conn) *User {
	//用户地址
	addr := conn.RemoteAddr().String()
	//实例化用户对象
	user := &User{
		Name: addr,
		Addr: addr,
		Conn: conn,
		Msg:  make(chan string),
	}

	//启动用户消息处理协程
	go user.handleMsg()
	return user
}
