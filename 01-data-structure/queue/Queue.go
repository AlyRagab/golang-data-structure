package main

import "fmt"

type Queue struct {
	item []int
}

func (q *Queue) enqueue(i int) {
	q.item = append(q.item, i)
}

func (q *Queue) dequeue() int {
	index, items := q.item[0], q.item[1:]
	q.item = items
	return index
}

func main() {
	q := Queue{}
	for i := 1; i <= 5; i++ {
		q.enqueue(i)
		fmt.Println("Enqueued element : ", i)
		fmt.Println("====")
		fmt.Println("Dequeued element : ", q.dequeue())
	}
}
