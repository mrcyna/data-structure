package structures

import "errors"

// Stack represents stack data structure implementation
type Stack struct {
	Items []interface{}
}

// NewStack returns a new Stack
func NewStack() Stack {
	return Stack{}
}

// Push add item to Stack
func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

// Pop remove item from Stack
func (s *Stack) Pop() (interface{}, error) {
	lastIndex := s.LastIndex()
	if lastIndex == -1 {
		return nil, errors.New("there is no item in the stack")
	}
	top := s.Items[s.LastIndex()]
	s.Items = s.Items[:s.LastIndex()]

	return top, nil
}

// LastIndex returns the last index of stack
func (s *Stack) LastIndex() int {
	return len(s.Items) - 1
}
