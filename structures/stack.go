package structures

import "errors"

type Stack struct {
	Items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	lastIndex := s.LastIndex()
	if lastIndex == -1 {
		return nil, errors.New("there is no item in the stack")
	}
	top := s.Items[s.LastIndex()]
	s.Items = s.Items[:s.LastIndex()]

	return top, nil
}

func (s *Stack) LastIndex() int {
	return len(s.Items) - 1
}
