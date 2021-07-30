package tests

import (
	"testing"

	"github.com/mrcyna/data-structures/structures"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &structures.BinarySearchTree{Key: 100}

	if tree.Key != 100 {
		t.Error("key should be 100")
	}

	tree.Insert(200)

	if tree.Left != nil {
		t.Error("tree left key should be nil")
	}
	if tree.Right.Key != 200 {
		t.Error("tree right key should be 200")
	}

	tree.Insert(300)

	if tree.Right.Left != nil {
		t.Error("tree right's left should be nil")
	}
	if tree.Right.Right.Key != 300 {
		t.Error("tree right's right should be 300")
	}

	if !tree.Search(100) {
		t.Error("100 should be exists in tree")
	}

	if !tree.Search(200) {
		t.Error("200 should be exists in tree")
	}

	if !tree.Search(300) {
		t.Error("300 should be exists in tree")
	}

	if tree.Search(10) {
		t.Error("10 should not exists in tree")
	}

	if tree.Search(50) {
		t.Error("50 should not exists in tree")
	}
}
