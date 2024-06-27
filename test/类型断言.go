package main

import (
	"container/list"
	"fmt"
)

type entryTest struct {
	key   int
	value int
}

func main() {

	// 初始化一个list.Element
	elem := &list.Element{
		Value: &entryTest{key: 1, value: 100},
	}

	// 类型断言
	e := elem.Value.(*entryTest)

	// 现在我们可以访问entry结构体中的key和value字段
	fmt.Println(e.key)   // 输出: 1
	fmt.Println(e.value) // 输出: 100

}
