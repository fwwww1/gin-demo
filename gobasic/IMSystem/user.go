package main

import (
	"net"
	"strings"
)

// 用户结构体

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	Server *Server
}

// 创建一个用户的API（新用户对象）

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		Server: server,
	}
	//启动监听当前user channel 的goroutine
	go user.ListenMessage()
	return user
}

// 用户上线

func (user *User) Online() {

	//将新用户添加到在线用户列表
	user.Server.mapLock.Lock()
	user.Server.OnlineMap[user.Name] = user
	user.Server.mapLock.Unlock()
	//广播当前用户上线消息
	user.Server.BroadCast(user, "已上线")
}

//用户下线

func (user *User) Offline() {
	//用户下线，将用户从OnlineMap中删除
	user.Server.mapLock.Lock()
	delete(user.Server.OnlineMap, user.Name)
	user.Server.mapLock.Unlock()
}

//给当前用户的客户端发送消息

func (user *User) SendMsg(onlineMsg string) {
	user.conn.Write([]byte(onlineMsg))
}

//用户处理消息

func (user *User) DoMessage(msg string) {
	if msg == "who" {
		//向用户发送在线消息
		for _, u := range user.Server.OnlineMap {
			onlineMsg := "[" + u.Addr + "]" + u.Name + ":" + "在线...\n"
			u.SendMsg(onlineMsg)
		}
		user.Server.BroadCast(user, msg)
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//获取新名字
		newName := strings.Split(msg, "|")[1]
		//判断新名字是否存在
		_, ok := user.Server.OnlineMap[newName]
		if ok {
			user.SendMsg("用户名已存在")
		} else {
			user.Server.mapLock.Lock()
			//删除旧名字
			delete(user.Server.OnlineMap, user.Name)
			//把用户绑定到以因名字为key的map
			user.Server.OnlineMap[newName] = user
			//将用户名字修改为新名字
			user.Name = newName
			user.Server.mapLock.Unlock()
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//1.获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("消息格式不正确，请使用\"to|name|msg\"格式")
			return
		}
		//2.根据用户名，得到对方的user对象
		u, ok := user.Server.OnlineMap[remoteName]
		if !ok {
			//用户名不存在
			user.SendMsg("用户不存在")
			return
		}
		//3.得到user对象后，获取要发送的消息，并发送给user
		toMsg := strings.Split(msg, "|")[2]
		if toMsg == "" {
			user.SendMsg("无消息内容请重发")
		}
		u.SendMsg(user.Name + "对您说：" + toMsg)
	} else {
		user.Server.BroadCast(user, msg)
	}

}

// 监听当前user中的channel，一旦有消息，立刻发送到客户端

func (user *User) ListenMessage() {
	for {
		message := <-user.C
		user.conn.Write([]byte(message + "\n"))
	}
}
