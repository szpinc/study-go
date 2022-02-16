package queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	q := Queue{
		Q:    []string{},
		Cond: &sync.Mutex{},
	}

	cha := make(chan struct{})

	go func() {
		q.Add("test")
		cha <- struct{}{}
	}()

	<-cha

	if item, ok := q.Get(); ok {
		fmt.Sprintln(item)
	}

}
