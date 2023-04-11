package queue

import (
	"fmt"
)

type Queue[T any] struct {
	elements []T
	size     int
}

func (q *Queue[T]) GetLength() int {
	return len(q.elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) GetElements() []T {
	return q.elements
}

func (q *Queue[T]) Enqueue(elem T) {
	if q.GetLength() == q.size {
		fmt.Println("Overflow")
		return
	}
	q.elements = append(q.elements, elem)
}

func (q *Queue[T]) Dequeue() T {
	var elem T

	if q.IsEmpty() {
		fmt.Println("Underflow")
		return elem
	}

	// Get first element from queue
	elem = q.elements[0]

	// Empty queue if it contains only 1 element
	if q.GetLength() == 1 {
		q.elements = nil
		return elem
	}

	// Pop first element of queue
	q.elements = q.elements[1:]

	return elem
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		// elements: make([]T, size),
		size: size,
	}
}
