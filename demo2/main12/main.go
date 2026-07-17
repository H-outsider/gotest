package main12

import "fmt"

func checkType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case bool:
		fmt.Printf("布尔值: %t\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}
