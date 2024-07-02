package main

import "fmt"

type Queue struct {
	data []int
	size int
}

func (q *Queue) GetLength() int {
	return len(q.data)
}

func (q *Queue) Enqueue(value int) {
	if q.GetLength() == q.size {
		panic("queue is full")
	}

	q.data = append(q.data, value)
	q.size++
}

func (q *Queue) Dequeue() int {
	if q.GetLength() == 0 {
		panic("queue is empty")
	}

	value := q.data[0]
	if q.GetLength() == 1 {
		q.data = nil
	} else {
		q.data = q.data[1:]
	}

	q.size--

	return value
}

func (q *Queue) Peek() int {
	if q.GetLength() == 0 {
		panic("queue is empty")
	}
	return q.data[0]

}

func main() {
	queue := Queue{size: 3}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("Queue length:", queue.GetLength())
	fmt.Println("Dequeued:", queue.Dequeue())
	fmt.Println("Queue length after Dequeued:", queue.GetLength())
	fmt.Println("Peeked:", queue.Peek())
	fmt.Println("Queue length after Peeked:", queue.GetLength())
}
