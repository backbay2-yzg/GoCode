package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}
type OnlineMsg struct {
	newMsg string
	Addr   string
}

var message = make(chan OnlineMsg)
var onlineMap map[string]Client

func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, clt := range onlineMap {
			if clt.Addr != msg.Addr {
				clt.C <- msg.newMsg
			}

		}
	}

}
func MakeMsg(clt Client, msg string) (message OnlineMsg) {
	var onlineMsg OnlineMsg
	onlineMsg.newMsg = "[" + clt.Addr + "]" + clt.Name + " say : " + msg
	onlineMsg.Addr = clt.Addr
	return onlineMsg
}
func HandlerConnection(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	clt := Client{make(chan string), "来自[" + addr + "]的访问：", addr}
	onlineMap[addr] = clt
	go WirteMsgToClient(conn, clt)
	message <- MakeMsg(clt, "login")
	hasData := make(chan bool)
	isQuit := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端: %s 退出\n", clt.Addr)
				return
			}
			if err != nil {
				fmt.Println("conn read err !")
				return
			}
			msg := string(buf[:n-2])
			if msg == "exit" && len(msg) == 4 {
				isQuit <- true
				fmt.Printf("客户端手动 %s 退出\n", clt.Addr)
				//delete(onlineMap, addr)
				//conn.Close()

				return
			}
			if msg == "who" && len(msg) == 3 {
				fmt.Println("online users list:\n")
				conn.Write([]byte("online users list:\n"))
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					fmt.Println(userInfo)
					conn.Write([]byte(userInfo))

				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				fmt.Println(newName)
				clt.Name = "来自[" + newName + "]的访问："
				onlineMap[addr] = clt
				conn.Write([]byte("rename successful\n"))
			} else {
				message <- MakeMsg(clt, msg)
			}
			hasData <- true
		}

	}()
	for {
		select {
		case <-isQuit:
			close(clt.C)
			delete(onlineMap, addr)
			message <- MakeMsg(clt, "logout")
			return
		case <-hasData:
		case <-time.After(time.Second * 15):
			close(clt.C)
			delete(onlineMap, addr)
			message <- MakeMsg(clt, "time out leave!!!")
			return
		}

	}
}

func WirteMsgToClient(conn net.Conn, clt Client) {
	for msg := range clt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net.Listen")
		return
	}
	defer listen.Close()
	go Manager()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept")
			return
		}
		go HandlerConnection(conn)
	}

}
