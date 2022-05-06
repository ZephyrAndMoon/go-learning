/*
 *@Title   Map
 *@Author  zephyr
 *@Update  2022-05-06
 */

package main

import "fmt"

/*
	创建：make(map[string]int)
	获取元素：m[key]
	key不存在时，获得 Value 类型的初始值
	用 value, ok := m[key] 来判断是否存在 key
	用 delete 删除一个 key

	! 遍历
	使用 range 遍历 key, 或者遍历 key, value 对
	不保证遍历顺序，如需顺序，需手动对 key 排序
	使用 len 获得元素个数

	! map 的 key
	map 使用哈希表, 必须可以比较相等
	除了 slice, map, function的内建类型都可以作为 key
	Struct 类型不包含上述字段，也可以作为 key
*/

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok1 := m["course"]
	fmt.Println(courseName, ok1)

	if carseName, ok2 := m["carse"]; ok2 {
		fmt.Println(carseName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
