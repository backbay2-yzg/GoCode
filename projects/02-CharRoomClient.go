package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net dial err: ", err)
		return
	}
	defer conn.Close()
	go func() {
		str := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read: ", err)
				continue
			}
			//fmt.Println(string(buf[:n]))
			_, err = conn.Write(str[:n])
			if err != nil {
				fmt.Println("conn.Write err: ", err)
				return
			}
		}
	}()
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}

}
