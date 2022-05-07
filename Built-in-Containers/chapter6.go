/*
 *@Title   寻找最长不含有重复字符的字串
 *@Author  zephyr
 *@Update  2022-05-06
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	rune 相当于 go 的 char
	- 使用 range 遍历 pos, rune 对
	- 使用 utf8.RuneCountInString 获得字符数量
	- 使用 len 获得字节长度
	- 使用 []byte 获得字节
*/

func main() {
	s := "YES我爱慕课网" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X\n", b)
	}
	for i, v := range s {
		fmt.Printf("(%d %X) ", i, v)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
