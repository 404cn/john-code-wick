package queue

type Queue struct {
	queue []interface{}
	len   int
}

func New() *Queue {
	q := &Queue{}
	q.queue = make([]interface{}, 0)
	q.len = 0
	return q
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Empty() bool {
	return q.len == 0
}

func (q *Queue) EnQueue(value interface{}) {
	arr := make([]interface{}, 1)
	arr[0] = value
	q.queue = append(q.queue, arr)
	q.len++
}

func (q *Queue) DeQueue() (value interface{}) {
	value, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return value
}
