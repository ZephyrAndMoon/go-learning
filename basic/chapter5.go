/*
 *@Title   循环
 *@Author  zephyr
 *@Update  2022-05-01
 */

/*
	for 的条件里不需要括号
	for 的条件里可以省略初始条件, 结束条件, 递增表达式
	for, if 后面条件没有括号
	if 条件里也可以定义变量
	没有 while
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBain(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBain(5),
		convertToBain(13),
		convertToBain(72683),
		convertToBain(0),
	)
	printFile("abc.txt")
}
