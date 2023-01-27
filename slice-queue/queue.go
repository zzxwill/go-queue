package queue

import "errors"

type Queue interface {
	Push(x interface{}) error
	Full() bool
	Pop() interface{}
	Len() int
}

// SliceQueue is a priority queue
type SliceQueue []interface{}

// Capacity is the capacity for each queue in Queues
const Capacity = 10

func (q *SliceQueue) Push(x interface{}) error {
	if q.Full() {
		return errors.New("queue is full")
	}
	*q = append(*q, x)
	return nil
}

func (q *SliceQueue) Full() bool {
	return q.Len() >= Capacity
}

func (q *SliceQueue) Pop() interface{} {
	if q == nil {
		return nil
	}
	old := *q
	n := old.Len()
	switch n {
	case 0:
		return nil
	case 1:
		*q = SliceQueue{}
		return old[0]
	default:
		x := old[0]
		*q = old[1:]
		return x
	}
}

func (q *SliceQueue) Len() int {
	return len(*q)
}
