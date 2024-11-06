package main

import (
	"fmt"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week2/tree"
)

func main() {
	root := bst.MakeNode(50)
	root.Insert(54)
	root.Insert(76)
	root.Insert(45)
	root.Insert(24)
	root.Insert(47)
	root.Insert(94)

	fmt.Println("InOrder (Sorted Order):")
	bst.InOrder(root)
	fmt.Println()
}
