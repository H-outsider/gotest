package main5

import (
	"encoding/json"
	"fmt"
)

type Speaker interface {
	Speak()
}

type Walker interface {
	Walk()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says woof!")
}

func (d Dog) Walk() {
	fmt.Println(d.Name, "is walking")
}

func describe(i interface{}) {
	// 尝试多种类型
	if s, ok := i.(string); ok {
		fmt.Printf("字符串: %q, 长度: %d\n", s, len(s))
		return
	}

	if n, ok := i.(int); ok {
		fmt.Printf("整数: %d, 平方: %d\n", n, n*n)
		return
	}

	if b, ok := i.(bool); ok {
		fmt.Printf("布尔: %v\n", b)
		return
	}

	fmt.Printf("未知类型: %T\n", i)
}

func processJSON(data string) {
	var result interface{}
	_ = json.Unmarshal([]byte(data), &result)

	// JSON对象解析为map[string]interface{}
	if m, ok := result.(map[string]interface{}); ok {
		fmt.Println("这是一个JSON对象:")
		for key, value := range m {
			processValue(key, value)
		}
		return
	}

	// JSON数组解析为[]interface{}
	if arr, ok := result.([]interface{}); ok {
		fmt.Println("这是一个JSON数组:")
		for i, v := range arr {
			fmt.Printf("  [%d]: %v (%T)\n", i, v, v)
		}
		return
	}

	fmt.Println("其他类型:", result)
}

func processValue(key string, value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("  %s: %q (字符串)\n", key, v)
	case float64: // JSON数字解析为float64
		fmt.Printf("  %s: %v (数字)\n", key, v)
	case bool:
		fmt.Printf("  %s: %v (布尔)\n", key, v)
	case nil:
		fmt.Printf("  %s: null\n", key)
	case []interface{}:
		fmt.Printf("  %s: (数组，%d个元素)\n", key, len(v))
	case map[string]interface{}:
		fmt.Printf("  %s: (嵌套对象)\n", key)
	default:
		fmt.Printf("  %s: %v (%T)\n", key, v, v)
	}
}
