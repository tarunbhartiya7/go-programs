package main

import "fmt"

type Queue struct {
	items []int
}

// pointer receiver functions
func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func main() {
	q := Queue{}
	q.Push(1)
	// q.Push(4)
	fmt.Println(q)

	q.Pop()
	fmt.Println(q)

	item, ok := q.Pop()
	if ok {
		fmt.Println("Popped item: ", item)
	} else {
		fmt.Println("Queue is empty")
	}

}
