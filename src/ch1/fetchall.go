package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 申请一个管道
	channel := make(chan string)
	for _, arg := range os.Args[1:] {
		// 异步执行url请求
		go fetch(arg, channel)
	}

	for range os.Args[1:] {
		// 接受管道消息
		fmt.Println(<-channel)
	}

	fmt.Printf("耗时:%.2fs", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	resp, err := http.Get(url) // 请求url
	if err != nil {
		ch <- fmt.Sprintf("fetch url error:%v", err)
		return
	}
	// io.Discard 无操作Writer，不记录数据，可以保证一定会成功
	body, err := io.Copy(io.Discard, resp.Body)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("read body error:%v", err)
		return
	}
	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", seconds, body, url)
}
