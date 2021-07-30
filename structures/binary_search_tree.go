package structures

// BinarySearchTree represents a node of binary search tree
type BinarySearchTree struct {
	Key   int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

// Insert insert new item into the BinarySearchTree
func (n *BinarySearchTree) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &BinarySearchTree{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &BinarySearchTree{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search check if the given value exists in BinarySearchTree
func (n *BinarySearchTree) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}
