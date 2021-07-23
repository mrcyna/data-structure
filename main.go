package main

import (
	"fmt"

	"github.com/mrcyna/data-structures/structures"
)

func main() {
	q := structures.NewQueue()
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	fmt.Println(q)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q)
}
