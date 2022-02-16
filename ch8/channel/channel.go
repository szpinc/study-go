package channel

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func mirrorQuery() string {
	resps := make(chan string, 3)

	go func() {
		if s, ok := request("http://asia.gopl.io"); ok {
			resps <- s
		}
	}()
	go func() {
		if s, ok := request("http://europe.gopl.io"); ok {
			resps <- s
		}
	}()
	go func() {
		if s, ok := request("http://americas.gopl.io"); ok {
			resps <- s
		}
	}()

	if content, ok := <-resps; ok {
		return content
	}
	return ""
}

func request(path string) (string, bool) {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
		return "", false
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", false
	}
	return fmt.Sprintf("%x", bytes), true
}
