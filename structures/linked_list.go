package structures

// Node represents a node of LinkedList
type Node struct {
	Data interface{}
	Next *Node
}

// LinkedList represent the LinkedList data structure
type LinkedList struct {
	Head   *Node
	Length int
}

func NewLinkedList() LinkedList {
	return LinkedList{
		Head:   nil,
		Length: 0,
	}
}

// Prepend prepend a Node to the head of LinkedList
func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

// Iterate iterate all Node of the LinkedList
func (l *LinkedList) Iterate(f func(n *Node)) {
	pointer := l.Head

	for pointer.Next != nil {
		f(pointer)
		pointer = pointer.Next
	}
	f(pointer)
}
