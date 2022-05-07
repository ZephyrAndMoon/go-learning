/*
 *@Title   面向对象
 *@Author  zephyr
 *@Update  2022-05-07
 */

package main

import "fmt"

/*
	go 语言仅支持封装，不支持继承和多态
	go 语言没有 class 只有 struct


	为结构定义方法
	func (node TreeNode) print(){
		fmt.Print(node.Value)
	}
	显示定义和命名方法接收者

	使用指针作为方法接收者
	func(node *TreeNode) setValue(value int){
		node.Value = Value
	}
	只有使用指针才可以改变结构内容
	nil 指针也可以调用方法


	要改变内容必须使用指针接收者
	结构过大也要考虑使用指针接收者
	一致性：如有指针接收者，最好都是指针接收者

	值接收者是 go 语言特有（其实 js 也是 ）
	值/指针接收者均可接收值/指针
*/

type treeNode struct {
	value       int
	left, right *treeNode
}

// 拷贝一份值进行打印
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 传递值的地址进行修改
func (node *treeNode) setValue(value int) {
	node.value = value
}

func createNode(value int) *treeNode {
	// 返回局部变量的地址
	return &treeNode{value: value}
}

func (node treeNode) traverse() {
	fmt.Println(node)
	// if node == nil {
	// 	return
	// }

	// node.left.traverse()
	// node.print()
	// node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	fmt.Println(root)

	root.left.right = createNode(2)

	root.print()
	root.right.left.setValue(4)
	root.right.left.print()
	root.traverse()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
