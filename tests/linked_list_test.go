package tests

import (
	"testing"

	"github.com/mrcyna/data-structures/structures"
)

func TestLinkedList(t *testing.T) {
	ll := structures.NewLinkedList()

	if ll.Length != 0 {
		t.Error("linked list should be zero size at the beginning")
	}
}

func TestPrependOneItemToTheLinkedList(t *testing.T) {
	ll := structures.NewLinkedList()

	ll.Prepend(&structures.Node{
		Data: 100,
		Next: nil,
	})

	if ll.Length != 1 {
		t.Error("linked list size should be one")
	}

	if ll.Head.Data != 100 {
		t.Error("linked list head data should be 100")
	}
}

func TestPrependThreeItemToTheLinkedList(t *testing.T) {
	ll := structures.NewLinkedList()

	ll.Prepend(&structures.Node{
		Data: 100,
		Next: nil,
	})

	ll.Prepend(&structures.Node{
		Data: 200,
		Next: nil,
	})

	ll.Prepend(&structures.Node{
		Data: 300,
		Next: nil,
	})

	if ll.Length != 3 {
		t.Error("linked list size should be three")
	}

	if ll.Head.Data != 300 {
		t.Error("linked list head data should be 300")
	}

	if ll.Head.Next.Data != 200 {
		t.Error("linked list head data should be 200")
	}

	if ll.Head.Next.Next.Data != 100 {
		t.Error("linked list head data should be 100")
	}

	sum := 0
	ll.Iterate(func(n *structures.Node) {
		sum += n.Data.(int)
	})
	if sum != 600 {
		t.Error("total sum of the nodes should be 600")
	}
}
