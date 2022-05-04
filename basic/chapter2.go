/*
 *@Title   内建变量类型
 *@Author  zephyr
 *@Update  2022-05-01
 */

/*
   内建变量类型
   bool, string
   (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
   byte, rune
   float32, float64, complex64, complex128
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
// 类型转换是强制的
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
