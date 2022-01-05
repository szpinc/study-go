package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	os.Args = append(os.Args, "http://www.baidu.com")

	for _, arg := range os.Args {
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
		fmt.Println(body)
	}

}
