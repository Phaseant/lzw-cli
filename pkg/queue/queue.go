package queue

// stolen from github.com/golang-collections/collections/blob/master/queue/queue.go

type (
	Queue struct {
		start, end *node
		length     int
	}
	node struct {
		value string
		next  *node
	}
)

func New() *Queue {
	return &Queue{nil, nil, 0}
}

func (q *Queue) Dequeue() string {
	if q.length == 0 {
		return "0"
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return n.value
}

func (q *Queue) Enqueue(value string) {
	n := &node{value, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

func (q *Queue) Len() int {
	return q.length
}

func (q *Queue) Peek() string {
	if q.length == 0 {
		return ""
	}
	return string(q.start.value)
}

func (q *Queue) Clean() {
	q.start = nil
	q.end = nil
	q.length = 0
}

func (q *Queue) String() string {
	all := ""
	cur := q.start
	for cur != nil {
		all += string(cur.value)
		cur = cur.next
	}
	return all
}
