package main

import (
	"fmt"

	"github.com/pratikms/simple-queue/queue"
)

func main() {
	q := queue.NewQueue[int](6)

	fmt.Println("en: ", q.GetLength())
	// Push to queue
	q.Enqueue(3)
	fmt.Println("len: ", q.GetLength())
	q.Enqueue(5)
	fmt.Println("len: ", q.GetLength())
	q.Enqueue(3)
	fmt.Println("len: ", q.GetLength())
	q.Enqueue(6)
	fmt.Println("len: ", q.GetLength())
	q.Enqueue(1)
	q.Enqueue(9)
	q.Enqueue(19)

	fmt.Printf("%+v", q.GetElements())
	q.Dequeue()
	q.Enqueue(4)
	fmt.Printf("%+v", q.GetElements())
}
