package main

type (
	BinaryNode struct {
		left  *BinaryNode
		right *BinaryNode
		data  int
	}

	BinaryTree struct {
		root *BinaryNode
	}
)

func (tree *BinaryTree) insert(data int) *BinaryTree {
	if tree.root == nil {
		node := &BinaryNode{
			left:  nil,
			right: nil,
			data:  data,
		}
		tree.root = node
	} else {
		tree.root.insert(data)
	}
	return tree
}

func (node *BinaryNode) insert(data int) {
	if node == nil {
		return
	} else if data <= node.data {
		if node.left == nil {
			node.left = &BinaryNode{
				left:  nil,
				right: nil,
				data:  data,
			}
		} else {
			node.left.insert(data)
		}
	} else {
		if node.right == nil {
			node.right = &BinaryNode{
				left:  nil,
				right: nil,
				data:  data,
			}
		} else {
			node.right.insert(data)
		}
	}
}
