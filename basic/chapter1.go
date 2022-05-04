/*
 *@Title   变量定义
 *@Author  zephyr
 *@Update  2022-05-01
 */

package main

import "fmt"

/*
	变量定义
	使用 var 关键字
	- var a,b,c bool
	- var s1, s2 string = "hello world"
	- 可放在函数内或者包内
	- 使用 var() 集中定义变量

	编译器自动决定类型
	使用 := 定义变量，只能在函数内使用
*/

// 函数外部使用定义不能用 ":="
// go 中没有全局变量的说法 是包内部变量
var aa = 3
var ss = "ss"

var (
	cc = 3
	qq = "qq"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, qq, ss)
}
