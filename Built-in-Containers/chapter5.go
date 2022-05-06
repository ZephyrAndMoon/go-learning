/*
 *@Title   寻找最长不含有重复字符的字串
 *@Author  zephyr
 *@Update  2022-05-06
 */

/*
	abcabcbb -> abc
	bbbbb -> b
	pwwkew -> wke

	对于每一个字母 x
	- lastOccurred[x]不存在，或者 < start 则无需操作
	- lastOccurred[x] >= start 则更新 start
	- 更新 lastOccurred[x], 更新 maxLength
*/

package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
}
