package structures

import "errors"

// Queue represents queue data structure implementation
type Queue struct {
	Items []interface{}
}

// NewQueue returns a new Queue
func NewQueue() Queue {
	return Queue{}
}

// Enqueue add item to Queue
func (q *Queue) Enqueue(item interface{}) {
	q.Items = append(q.Items, item)
}

// Dequeue remove item from Queue
func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.Items) == 0 {
		return nil, errors.New("there is no item in queue")
	}

	top := q.Items[0]
	q.Items = q.Items[1:]

	return top, nil
}
