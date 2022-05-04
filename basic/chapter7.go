/*
 *@Title   指针
 *@Author  zephyr
 *@Update  2022-05-01
 */

/*
	不能运算
*/

package main

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
