package main5

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var i interface{} = "hello"

	// 断言为string
	s := i.(string)
	fmt.Println("字符串:", s)
	fmt.Println("长度:", len(s))

	// 断言错误类型会panic
	// n := i.(int)  // panic: interface conversion
}

func TestMain2(t *testing.T) {
	var i interface{} = 42

	// 安全断言
	s, ok := i.(string)
	if ok {
		fmt.Println("是字符串:", s)
	} else {
		fmt.Println("不是字符串，s的零值:", s)
	}

	// 正确的类型
	n, ok := i.(int)
	if ok {
		fmt.Println("是整数:", n)
	}
}

func TestMain3(t *testing.T) {
	var s Speaker = Dog{Name: "Buddy"}

	// 断言为另一个接口
	if w, ok := s.(Walker); ok {
		fmt.Println("Dog也是Walker:")
		w.Walk()
	}

	// 断言为具体类型
	if dog, ok := s.(Dog); ok {
		fmt.Println("Dog的名字:", dog.Name)
	}
}

func TestMain4(t *testing.T) {
	describe("hello")
	describe(42)
	describe(true)
	describe(3.14)
}

func TestMain5(t *testing.T) {
	jsonObj := `{
		"name": "张三",
		"age": 25,
		"active": true,
		"score": null,
		"tags": ["go", "python"]
	}`

	processJSON(jsonObj)

	fmt.Println()

	jsonArr := `[1, 2, 3, "hello"]`
	processJSON(jsonArr)
}
