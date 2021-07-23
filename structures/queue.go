package structures

type Queue struct {
	Items []interface{}
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Dequeue() interface{} {
	top := q.Items[0]
	q.Items = q.Items[1:]

	return top
}
