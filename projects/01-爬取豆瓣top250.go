package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var start, end int
	fmt.Println("请输入爬取的起始页：(页数>=1)")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的终止页：(页数>=起始页)")
	fmt.Scan(&end)
	workByDouban(start, end)

}

func workByDouban(start int, end int) {
	page := make(chan int)
	fmt.Printf("正在爬取豆瓣第%d页到%d页...\n", start, end)
	for i := start; i <= end; i++ {
		go spiderPageByDouban(i, page)
		//spiderPageByDouban(i,page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页豆瓣Top250数据导入成功\n", <-page)
	}

}
func Submatch(s string, match string) (str []string) {
	reg := regexp.MustCompile(s)
	result := reg.FindAllStringSubmatch(match, -1)
	for _, value := range result {
		str = append(str, value[1])
	}
	return
}
func spiderPageByDouban(i int, page chan int) {

	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	//url := "https://tieba.baidu.com/f?kw=%E6%B3%89%E5%B7%9E%E5%B8%88%E8%8C%83%E5%AD%A6%E9%99%A2&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)

	result, err := HttpGetByDouban(url)
	fmt.Println("--------")
	//fmt.Println(result)

	if err != nil {
		fmt.Println("HttpGetByDouban: ", err)
		return
	}
	//reg1 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	//filmName := reg1.FindAllStringSubmatch(result, -1)
	//reg2 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	//stars := reg2.FindAllStringSubmatch(result, -1)
	//reg3 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	//nums := reg3.FindAllStringSubmatch(result, -1)

	filmName := Submatch(`<img width="100" alt="(.*?)"`, result)
	fmt.Println(filmName)

	stars := Submatch(`<span class="rating_num" property="v:average">(.*?)</span>`, result)
	fmt.Println(stars)
	//
	nums := Submatch(`<span>(.*?)人评价</span>`, result)
	fmt.Println(nums)
	SaveFile(i, filmName, stars, nums)
	page <- i

}

func SaveFile(index int, filmName []string, stars []string, nums []string) {

	f, err := os.Create("D:/第" + strconv.Itoa(index) + "页豆瓣Top250数据.txt")
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}
	n := len(filmName)
	f.WriteString("电影名\t\t\t评分\t\t\t评分人数\t\t\t" + "\n")
	for i := 0; i < n; i++ {
		f.WriteString(filmName[i] + "\t\t\t" + stars[i] + "\t\t\t" + nums[i] + "\t\t\t" + "\n")
	}
	f.Close()

	//if n == 0 {
	//	fmt.Println("read finish...")
	//	return
	//}
	//if err != nil && err != io.EOF {
	//	fmt.Println("f.Read err : ", err)
	//	return
	//}
	//defer f.Close()
}

var header = map[string]string{
	"Host":                      "movie.douban.com",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Referer":                   "https://movie.douban.com/top250",
}

func HttpGetByDouban(url string) (result string, err error) {
	fmt.Println(url)
	client := &http.Client{}

	req, err1 := http.NewRequest("GET", url, nil)
	for key, value := range header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	//resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}

	return
}
