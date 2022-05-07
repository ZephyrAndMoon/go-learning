package main

import (
	"fmt"

	"../../tree"
)

func CreateNode(value int) *tree.TreeNode {
	// 返回局部变量的地址
	return &tree.TreeNode{Value: value}
}

func main() {
	var root tree.TreeNode = tree.TreeNode{Value: 3}

	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, {}, nil}
	root.Right.Left = new(tree.TreeNode)
	fmt.Println(root)

	root.Left.Right = CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	root.Traverse()

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
