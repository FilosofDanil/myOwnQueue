package main

import (
	"fmt"
	"myOwnQueue/pkg"
)

func main() {
	var q = pkg.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	for {
		if value, err := q.Dequeue(); err == nil {
			fmt.Println(value)
		} else {
			break
		}
	}
	fmt.Println(q.Dequeue())
}
