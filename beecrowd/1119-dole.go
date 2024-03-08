package main

import "errors"

type CircularQueue struct {
	length int
	items  []int
}

func (q *CircularQueue) Enqueue(item int) error {
	if len(q.items) == q.length {
		return errors.New("Queue is full")
	}
	q.items = append([]int{item}, q.items...)
	return nil
}

func (q *PQueue) Dequeue() (Cashier, error) {
	if len(q.items) == 0 {
		return Cashier{}, errors.New("Queue is empty")
	}
	i := q.items[0]
	q.items = q.items[1:]
	return i, nil
}

func (q *PQueue) Peek() (Cashier, error) {
	if len(q.items) == 0 {
		return Cashier{}, errors.New("Queue is empty")
	}
	return q.items[0], nil
}
