package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "read body error: %v\n", err)
			continue
		}
		fmt.Println(string(body))

		// 也可以直接用io.Copy来复制响应体到终端
		// io.Copy(os.Stdout, resp.Body)
	}

}
