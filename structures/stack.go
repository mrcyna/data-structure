package structures

type Stack struct {
	Items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() interface{} {
	top := s.Items[s.LastIndex()]
	s.Items = s.Items[:s.LastIndex()]

	return top
}

func (s *Stack) LastIndex() int {
	return len(s.Items) - 1
}
