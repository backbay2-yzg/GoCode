package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端请求：", r.URL)
	OpenAndSendFile(r.URL.String(), w)
}

func OpenAndSendFile(url string, w http.ResponseWriter) {
	pathFileByUrl := "G:/photo" + url
	f, err := os.Open(pathFileByUrl)

	if err != nil {
		w.Write([]byte("文件访问失败"))
		fmt.Println("os.Open err ")
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {

			if err == io.EOF {
				fmt.Println("文件读取完成。")
			} else {
				fmt.Println("f.Read err: ", err)
			}
			return
		}

		_, err = w.Write(buf[:n])
		if err != nil {
			fmt.Println("w.Write err: ", err)
			return
		}

	}
	//fmt.Println("成功")

}
func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8006", nil)
}
