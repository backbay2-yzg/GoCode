package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var start, end int
	fmt.Println("请输入爬取的起始页：(页数>=1)")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页：(页数>=起始页)")
	fmt.Scan(&end)
	working(start, end)

}

func HttpGet(url string) (result string, err error) {
	fmt.Println(url)
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	for {
		buf := make([]byte, 4096)
		n, err2 := resp.Body.Read(buf)
		defer resp.Body.Close()
		if err2 != nil && err2 != io.EOF {
			fmt.Println("resp.Bodu.Read err :", err2)
			return
		}
		if n == 0 {
			fmt.Println("网页读取完成！")
			break
		}
		result += string(buf[:n])
	}

	return

}
func SpriderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E6%B3%89%E5%B7%9E%E5%B8%88%E8%8C%83%E5%AD%A6%E9%99%A2&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	//url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err: ", err)
		return
	}
	//fmt.Printf("result %d-->%v\n",i,result)
	f, err := os.Create("第" + strconv.Itoa(i) + "页数据.html")
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}
func working(start int, end int) {
	fmt.Printf("正在爬取第%d到%d页网页中，请稍后...", start, end)

	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpriderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页数据完成！", <-page)
	}
}
