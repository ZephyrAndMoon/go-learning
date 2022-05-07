package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 拷贝一份值进行打印
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

// 传递值的地址进行修改
func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
