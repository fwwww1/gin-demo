package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 创建一个server类型

type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播的channel
	Message chan string
}

//监听Message channel，一旦有消息，就发送到所有的在线用户

func (Server *Server) ListenMessage() {
	for {
		msg := <-Server.Message
		Server.mapLock.Lock()
		for _, u := range Server.OnlineMap {
			u.C <- msg
		}
		Server.mapLock.Unlock()
	}
}

// 广播消息的方法(将消息传入Message channel)

func (Server *Server) BroadCast(user *User, msg string) {
	sendMessage := "[" + user.Addr + "]" + user.Name + ":" + msg
	Server.Message <- sendMessage
}

// 处理函数

func (Server *Server) Handler(conn net.Conn) {
	//一个新用户建立连接
	newUser := NewUser(conn, Server)
	newUser.Online()
	//监听用户是否活跃的管道
	isLive := make(chan bool)
	//读取用户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if n == 0 {
			newUser.Offline()
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("conn Read err:", err)
			return
		}
		//获取用户消息
		s := string(buf[:n-1])
		//广播消息
		newUser.DoMessage(s)
		//用户的任意消息，代表当前用户是活跃的
		isLive <- true
	}()
	//当前handler阻塞
	for {
		select {
		case <-isLive:
		//当前用户是活跃的，应重置定时器
		//不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(10 * time.Second):
			//已经超时
			//将当前的user关闭
			newUser.SendMsg("你被踢了")
			//销毁用的资源
			close(newUser.C)
			//关闭连接
			conn.Close()
			//退出handler
			return
		}
	}

}

//创建一个server接口(返回server对象)

func NewServer(Ip string, Port int) *Server {
	server := &Server{
		Ip:        Ip,
		Port:      Port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//启动服务器的接口

func (Server *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Server.Ip, Server.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	//close socket listen
	defer listener.Close()
	//启动监听Message channel的goroutine
	go Server.ListenMessage()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//do handler
		go Server.Handler(conn)
	}

}
