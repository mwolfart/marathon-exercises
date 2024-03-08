package main

import (
	"errors"
	"fmt"
)

type Cashier struct {
	id       int
	cap      int
	priority int
}

type PQueue struct {
	items []Cashier
}

func isPriorityLower(a Cashier, b Cashier) bool {
	return a.priority < b.priority
}

func isPrioritySameAndIdLower(a Cashier, b Cashier) bool {
	return a.priority == b.priority && a.id < b.id
}

func (q *PQueue) Enqueue(item Cashier) {
	i := 0
	for i < len(q.items) && (isPriorityLower(q.items[i], item) || isPrioritySameAndIdLower(q.items[i], item)) {
		i++
	}
	if i == len(q.items) {
		q.items = append(q.items, item)
	} else {
		q.items = append(q.items[:i+1], q.items[i:]...)
		q.items[i] = item
	}
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

func (q *PQueue) PeekLast() (Cashier, error) {
	if len(q.items) == 0 {
		return Cashier{}, errors.New("Queue is empty")
	}
	return q.items[len(q.items)-1], nil
}

func getTotalTime(cashiers *PQueue, customers []int) (int, error) {
	for _, c := range customers {
		x, _ := cashiers.Dequeue()
		priority := (x.cap * c) + x.priority
		cashiers.Enqueue(Cashier{x.id, x.cap, priority})
	}
	last, err := cashiers.PeekLast()
	if err != nil {
		return 0, errors.New("Unknown error")
	}
	return last.priority, nil
}

func main() {
	var numCustomers, numCashiers int
	fmt.Scanf("%d %d", &numCashiers, &numCustomers)

	cashiers := make([]Cashier, numCashiers)
	customers := make([]int, numCustomers)

	for i := 0; i < numCashiers; i++ {
		var cap int
		fmt.Scanf("%d", &cap)
		cashiers[i] = Cashier{i, cap, 0}
	}

	for i := 0; i < numCustomers; i++ {
		fmt.Scanf("%d", &customers[i])
	}

	total, err := getTotalTime(&PQueue{cashiers}, customers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)
}
