package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	Conn       net.Conn
	flag       int //当前client的模式
}

//创建client接口

func NewClient(ServerIp string, ServerPort int) (client *Client) {
	//创建客户端对象
	client = &Client{
		ServerIP:   ServerIp,
		ServerPort: ServerPort,
		flag:       999,
	}
	//连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerIp, ServerPort))
	if err != nil {
		fmt.Println("net.Dial err", err)
		return nil
	}

	client.Conn = conn

	//返回对象
	return client
}

// 处理server回应的消息，直接显示到标准输出即可

func (client *Client) DealResponse() {
	//一旦client.conn有数据，就直接copy数据到stdout上，永久阻塞监听
	io.Copy(os.Stdout, client.Conn)
}

// 获取用户输入的模式
func (client *Client) menu() bool {
	var flag int
	fmt.Println("1:公聊模式")
	fmt.Println("2:私聊模式")
	fmt.Println("3:更改用户名")
	fmt.Println("0:退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入合法范围内的数字")
		return false
	}
}

//查询在线用户

func (client *Client) SelectUser() {
	sendMsg := "who\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
}

// 私聊模式

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	//查询当前在线用户
	client.SelectUser()
	//输入聊天对象
	fmt.Println("请输入聊天对象[用户名]，exit退出")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		//输入聊天内容
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
		for len(chatMsg) != 0 {
			if len(chatMsg) != 0 {
				//把消息发送到服务器
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.Conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write err", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
		fmt.Println("请输入聊天对象[用户名]，exit退出")
		fmt.Scanln(&remoteName)
	}
}

// 公聊模式

func (client *Client) PublicChat() {
	var chatMsg string
	//提示用户输入消息
	fmt.Println("请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		//将消息发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err", err)
				break
			}
			chatMsg = ""
			fmt.Println("请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
	}
}

// 更改用户名

func (client *Client) UpdateName() bool {
	fmt.Println("请输入用户名")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return false
	}
	return true
}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			//公聊模式
			fmt.Println("公聊模式")
			client.PublicChat()
			break
		case 2:
			//私聊模式
			fmt.Println("私聊模式")
			client.PrivateChat()
			break
		case 3:
			//更改用户名
			client.UpdateName()
			fmt.Println("更改用户名模式")
			break
		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8888）")
}
func main() {
	//命令行解析
	flag.Parse()
	newClient := NewClient(serverIp, serverPort)
	if newClient == nil {
		fmt.Println("连接服务器失败")
		return
	}
	//开启一个goroutine去处理server的回执消息
	go newClient.DealResponse()
	fmt.Println("连接服务器成功")
	//启动客户端业务
	newClient.Run()
	select {}
}
