package netcat3

import (
	"io"
	"log"
	"net"
	"os"
	"testing"
)

// 测试channel同步
func TestSync(t *testing.T) {

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
		return
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Print("done")
		done <- struct{}{}
	}()

	io.Copy(conn, os.Stdin)
	conn.Close()
	<-done
}
