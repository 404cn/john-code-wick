package stack

type Stack struct {
	stack []interface{}
	len   int
}

func New() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0
	return s
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Push(value interface{}) {
	arr := make([]interface{}, 1)
	arr[0] = value
	s.stack = append(arr, s.stack...)
	s.len++
}

func (s *Stack) Peek() interface{} {
	return s.stack[0]
}

func (s *Stack) Pop() (value interface{}) {
	value, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}
