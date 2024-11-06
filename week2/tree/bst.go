package bst

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func MakeNode(num int) *Node{
	return &Node{Value: num}
}

func (tree *Node) Insert(num int){
	if(tree == nil){
		MakeNode(num)
	}else if tree.Value > num {
		tree.Right = MakeNode(num)
	}else{
		tree.Left = MakeNode(num)
	}
}

func InOrder(n *Node){
	if(n != nil){
		InOrder(n.Left)
		fmt.Print(n.Value, " ")
		InOrder(n.Right)
	}
}