/*
 *@Title   数组 Array
 *@Author  zephyr
 *@Update  2022-05-05
 */

package main

import "fmt"

/*
	为什么要用 range
	- 意义明确，美观
	- C++ 没有类似能力
	- Java/Py 只能 for each value, 不能同时获取 i, v
*/

/*
	数组是值类型
	[10]int 和 20[int] 是不同类型
	调用 func f(arr [10]int) 会拷贝数组
	在 go 语言中一般不直接使用数组
*/

func printArray(arr [5]int) {
	// 不会影响调用时传入的参数
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	// cannot use arr2 (variable of type [3]int) as [5]int value in argument to printArray
	// printArray(arr2)

	printArray(arr1)
	printArray(arr3)
	fmt.Println(arr1, arr3) // 还是 [0 0 0 0 0] / [2 4 6 8 10]
}
