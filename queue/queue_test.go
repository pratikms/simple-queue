package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue[int](5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	err := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()

		q.Enqueue(6)

		return nil
	}()

	if err != nil {
		t.Errorf("Unexpected error on Enqueue operation on a full queue, but got none")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int](5)
	q.Enqueue(1)
	q.Enqueue(2)

	elem := q.Dequeue()

	if elem != 1 {
		t.Errorf("Expected 1 on Dequeue, but got %v", elem)
	}

	elem = q.Dequeue()

	if elem != 2 {
		t.Errorf("Expected 2 on Dequeue, but got %v", elem)
	}

	err := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = r.(error)
			}
		}()

		q.Dequeue()

		return nil
	}()

	if err != nil {
		t.Errorf("Unxpected error on Dequeue operation on an empty queue, but got none")
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue[int](5)

	if !q.IsEmpty() {
		t.Errorf("Expected true on IsEmpty for a new queue, but got false")
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Errorf("Expected false on IsEmpty for a queue with one element, but got true")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("Expected true on IsEmpty for an empty queue, but got false")
	}
}

func TestGetLength(t *testing.T) {
	q := NewQueue[int](5)

	if q.GetLength() != 0 {
		t.Errorf("Expected 0 for GetLength for a new queue, but got %v", q.GetLength())
	}

	q.Enqueue(1)

	if q.GetLength() != 1 {
		t.Errorf("Expected 1 for GetLength for a queue with one element, but got %v", q.GetLength())
	}

	q.Enqueue(2)

	if q.GetLength() != 2 {
		t.Errorf("Expected 2 for GetLength for a queue with two elements, but got %v", q.GetLength())
	}
}

func TestGetElements(t *testing.T) {
	q := NewQueue[int](5)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	elements := q.GetElements()

	if len(elements) != 3 {
		t.Errorf("Expected 3 elements for GetElements, but got %v", len(elements))
	}

	if elements[0] != 1 || elements[1] != 2 || elements[2] != 3 {
		t.Errorf("Unexpected elements in GetElements result: %v", elements)
	}
}
