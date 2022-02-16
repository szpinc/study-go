package queue

import (
	"sync"
)

type Queue struct {
	Q    []string
	Cond sync.Locker
}

func (q *Queue) Add(str string) {
	q.Cond.Lock()
	defer q.Cond.Unlock()
	q.Q = append(q.Q, str)
}

func (q *Queue) Get() (string, bool) {
	q.Cond.Lock()
	defer q.Cond.Unlock()

	var item string

	if len(q.Q) == 0 {
		return item, false
	}
	item, q.Q = q.Q[0], q.Q[1:]
	return item, true
}
