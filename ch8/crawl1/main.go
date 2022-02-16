package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	channelExtract()
}

func channelExtract() {

	workList := make(chan []string)
	unseenLinks := make(chan string)

	go func() { workList <- os.Args[1:] }()

	// 起20个常驻groutine执行
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if findLink, err := Extract(link); err == nil {
					// 这里用新的groutine处理是防止死锁
					go func() { workList <- findLink }()
				}
			}
		}()
	}

	seenLinkMap := make(map[string]bool)

	for list := range workList {
		for _, link := range list {
			if !seenLinkMap[link] {
				seenLinkMap[link] = true
				write(link)
				unseenLinks <- link
			}
		}
	}

}

func write(link string) {

	u := url.URL{}
	url, err := u.Parse(link)
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp, err := http.Get(link)

	if err != nil {
		log.Fatal(err)
		return
	}

	if resp == nil {
		log.Fatal("repsonse body is nil")
		return
	}

	defer resp.Body.Close()

	filePath := "/Users/szp/Desktop/books.studygolang.com" + url.RequestURI()

	log.Println("file path:", filePath)

	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	io.Copy(file, resp.Body)
}

func Extract(url string) ([]string, error) {

	log.Printf("request url:%s \n", url)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body := resp.Body

	if body == nil {
		return nil, errors.New("body is nil")
	}

	defer body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get url:%s error:%d", url, resp.StatusCode)
	}

	contentTypes, ok := resp.Header["Content-Type"]

	if !ok {
		return nil, fmt.Errorf("requset[%s] is not text", url)
	}

	responseText := false

	for _, contentType := range contentTypes {
		if strings.Contains(contentType, "text/html") {
			responseText = true
			break
		}
	}

	if !responseText {
		return nil, fmt.Errorf("requset[%s] is not text", url)
	}

	log.Println("is text response")
	// 解析html
	doc, err := html.Parse(body)

	if err != nil {
		return nil, err
	}

	var urls []string

	visitNodeFunc := func(node *html.Node) {
		if node.Type != html.ElementNode || node.Data != "a" {
			return
		}
		for _, a := range node.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := resp.Request.URL.Parse(a.Val)
			if err != nil {
				fmt.Println(err)
				continue
			}
			urls = append(urls, link.String())
		}
	}
	forEachNode(doc, visitNodeFunc, nil)
	return urls, nil
}

// forEachNode 遍历节点
func forEachNode(node *html.Node, pre, post func(n *html.Node)) {
	// 前置处理
	if pre != nil {
		pre(node)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		forEachNode(n, pre, post)
	}
	// 后置处理
	if post != nil {
		post(node)
	}
}
