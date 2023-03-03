package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// 存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	//reImg = `^https://th\.wallhaven\.cc/[^\s]+\.jpg$`
	//reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	//初始化管道
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 100)
	//2.爬虫协程
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go getImages("https://wallhaven.cc/latest?page=" + strconv.Itoa(i+1))
	}
	//3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		//go DownloadImg()
	}
	waitGroup.Wait()
}

func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 10 {
			close(chanImageUrls)
			break
		}
	}
	fmt.Println("count", count)
	waitGroup.Done()
}

func getImages(url string) {
	resp, err := http.Get(url)
	HandleError(err, "http get url error")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	HandleError(err, "io read err")
	doc.Find("#thumbs img").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("data-src")
		r := strings.NewReplacer("th", "w", "small", "full")
		all := r.Replace(url)
		index := strings.LastIndex(all, "/")
		ss := all[:index] + "/wallhaven-" + all[index+1:]
		fmt.Println(ss)
		chanImageUrls <- ss
	})
	chanTask <- url
	waitGroup.Done()
	return
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// DownloadImg 下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// GetFilenameFromUrl 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// DownloadFile 下载图片
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.get.url")
	defer resp.Body.Close()
	//bytes, err := ioutil.ReadAll(resp.Body)
	bytes, err := io.ReadAll(resp.Body)
	HandleError(err, "resp.body")
	filename = "/Users/macbook/Downloads/pi/" + filename
	// 写出数据
	//err = ioutil.WriteFile(filename, bytes, 0666)
	err = os.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}
