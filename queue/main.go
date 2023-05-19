package main

import "fmt"

type queue struct {
	queue []int
}

func (q *queue) enqueue(data int) {
	q.queue = append(q.queue, data)
}

func (q *queue) dequeue() *int {
	if len(q.queue) > 0 {
		first := q.queue[0]
		q.queue = append(make([]int, 0), q.queue[1:]...)
		return &first
	}
	return nil
}

func (q *queue) reverse() {
	queueLength := len(q.queue)
	newQueue := make([]int, queueLength)
	for i := 0; i < queueLength; i++ {
		newQueue[queueLength-i-1] = *q.dequeue()
	}
	q.queue = newQueue
}

func main() {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	fmt.Println(q.queue)
	q.reverse()
	fmt.Println(q.queue)
	fmt.Println(*q.dequeue())
	fmt.Println(*q.dequeue())
	fmt.Println(*q.dequeue())
	fmt.Println(*q.dequeue())
}
