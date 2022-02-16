package pipeline

import "testing"

func TestPipeline(t *testing.T) {

	n := make(chan int)

	s := make(chan int)

	go func() {
		for i := 0; ; i++ {
			n <- i
		}
	}()

	go func() {
		for {
			x := <-n
			s <- x * x
		}
	}()

	for {
		println(<-s)
	}

}

// 改良版
// 使用单方向通道
func TestPipeline2(t *testing.T) {

	n := make(chan int)
	s := make(chan int)

	go counter(n)
	go squarer(s, n)
	printer(s)
}

// 累加
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		println(x)
	}
}
