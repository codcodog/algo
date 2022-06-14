package data_structure

import "fmt"

type Queue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

func NewQueue(c int) *Queue {
	return &Queue{
		data:     make([]interface{}, c),
		capacity: c,
		head:     0,
		tail:     0}
}

func (q *Queue) Enqueue(v interface{}) bool {
	if q.tail == q.capacity {
		return false
	}

	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *Queue) Dequeue() interface{} {
	if q.head == q.tail {
		return nil
	}

	v := q.data[q.head]
	q.head++
	return v
}

func (q *Queue) Print() {
	if q.head == q.tail {
		fmt.Println("empty queue")
		return
	}

	for i := q.head; i <= q.tail-1; i++ {
		fmt.Printf("%v ", q.data[i])
	}
	fmt.Println()
}
