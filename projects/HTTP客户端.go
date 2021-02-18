package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://127.0.0.1:8006/logo.png")
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println("http.Get err ")
		return
	}
	defer resp.Body.Close()

	fmt.Println("Header: ", resp.Header)
	fmt.Println("Status: ", resp.Status)
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("Proto: ", resp.Proto)
	var result string
	buf := make([]byte, 1024*8)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("--------read finnish")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err: ", err)
			return
		}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)
}
