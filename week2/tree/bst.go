package bst

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func MakeNode(num int) *Node {
	return &Node{Value: num}
}

func (tree *Node) Insert(num int) {
	if tree == nil {
		tree = MakeNode(num)
	} else if num < tree.Value {
		if tree.Left == nil {
			tree.Left = MakeNode(num) 
		} else {
			tree.Left.Insert(num)
		}
	} else {
		if tree.Right == nil {
			tree.Right = MakeNode(num)
		} else {
			tree.Right.Insert(num)
		}
	}
}

func InOrder(n *Node) {
	if n != nil {
		InOrder(n.Left)
		fmt.Print(n.Value, " ")
		InOrder(n.Right)
	}
}