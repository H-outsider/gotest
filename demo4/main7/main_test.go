package main7

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var i interface{}

	i = 42
	fmt.Printf("int: %v, type: %T\n", i, i)

	i = "hello"
	fmt.Printf("string: %v, type: %T\n", i, i)

	i = true
	fmt.Printf("bool: %v, type: %T\n", i, i)

	i = []int{1, 2, 3}
	fmt.Printf("slice: %v, type: %T\n", i, i)

	i = map[string]int{"a": 1}
	fmt.Printf("map: %v, type: %T\n", i, i)
}

func TestMain2(t *testing.T) {
	printValue(42)
	printValue("hello")
	printValue([]int{1, 2, 3})

	fmt.Println("---")
	printAll(1, "two", 3.0, true)
}

func TestMain3(t *testing.T) {
	// 存储不同类型的切片
	mixed := []any{
		42,
		"hello",
		true,
		3.14,
		[]int{1, 2, 3},
	}

	for _, v := range mixed {
		switch val := v.(type) {
		case int:
			fmt.Println("整数:", val*2)
		case string:
			fmt.Println("字符串长度:", len(val))
		case bool:
			fmt.Println("布尔:", val)
		case float64:
			fmt.Println("浮点:", val)
		default:
			fmt.Println("其他:", val)
		}
	}
}

func TestMain4(t *testing.T) {
	// 类似动态语言的对象
	person := map[string]any{
		"name":   "张三",
		"age":    25,
		"active": true,
		"scores": []int{90, 85, 92},
	}

	fmt.Println("姓名:", person["name"])
	fmt.Println("年龄:", person["age"])

	// 使用时需要类型断言
	if scores, ok := person["scores"].([]int); ok {
		total := 0
		for _, s := range scores {
			total += s
		}
		fmt.Println("总分:", total)
	}
}

func TestMain5(t *testing.T) {
	// 空接口版本
	fmt.Println("any:", maxAny(3, 5))
	fmt.Println("any:", maxAny(3.14, 2.71))

	// 泛型版本
	fmt.Println("generic:", maxGeneric(3, 5))
	fmt.Println("generic:", maxGeneric(3.14, 2.71))
}

func TestMain6(t *testing.T) {
	config := NewConfig()

	config.Set("host", "localhost")
	config.Set("port", 8080)
	config.Set("debug", true)
	config.Set("timeout", 30.5)

	fmt.Println("host:", config.GetString("host"))
	fmt.Println("port:", config.GetInt("port"))
	fmt.Println("debug:", config.GetBool("debug"))
	fmt.Println("timeout as string:", config.GetString("timeout"))
	fmt.Println("timeout as int:", config.GetInt("timeout"))
}
