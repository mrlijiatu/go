package main

import "fmt"

// map 可以理解为其他语言的字典，或者哈希表，以kv的形式进行数据存贮，与其他语言不同的是k可以是多类型的（string int，等
// v的值也是多类型的）
// 定义map的三种方式
func main() {
	var m map[string]string
	m = map[string]string{}
	m["name"] = "xiaoma"
	m["sex"] = "nan"

	m1 := map[string]interface{}{}
	m1["class1"] = "a1"
	m1["ui"] = []int{1, 2, 3, 4, 5}
	m1["yi"] = true
	m1["oi"] = 80

	m2 := make(map[int]string)
	m2[80] = "yuwen"
	delete(m1, "yi")
	//取map内的值使用for range
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println(m1)
}
