/*
 *@Title   常量与枚举
 *@Author  zephyr
 *@Update  2022-05-01
 */
package main

import (
	"fmt"
	"math"
)

/*
	const filename = 'abc.txt'
	const 数值可作为各种类型使用
*/

func constants() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt((a*a + b*b)))
	fmt.Println(filename, c)
}

func enums() {

	// const (
	// 	cpp    = 1
	// 	java   = 2
	// 	python = 3
	// 	golang = 4
	// )
	const (
		cpp = iota // 表示是自增值
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	constants()
	enums()
}
