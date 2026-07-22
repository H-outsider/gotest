package main6

import (
	"fmt"
	"testing"
	"time"
)

func TestMain1(t *testing.T) {
	describeValue(42)
	describeValue("hello")
	describeValue(true)
	describeValue(3.14)
	describeValue([]int{1, 2, 3})
}

func TestMain2(t *testing.T) {
	printNumber(42)
	printNumber(uint(100))
	printNumber(3.14)
	printNumber("hello")
}

func TestMain3(t *testing.T) {
	process(Dog{Name: "Buddy"}) // Dog同时实现了两个，匹配第一个
	process(Cat{Name: "Kitty"}) // Cat只实现了Speaker
	process("hello")
}

func TestMain4(t *testing.T) {
	fmt.Println(getType(42))
	fmt.Println(getType("hello"))
	fmt.Println(getType(true))
}

func TestMain5(t *testing.T) {
	handleNil(nil)
	handleNil(42)

	var p *int
	handleNil(p) // 不是nil case，是*int类型
}

func TestMain6(t *testing.T) {
	values := []interface{}{
		nil,
		true,
		false,
		42,
		3.14159,
		"Hello",
		"",
		time.Now(),
		[]string{"Go", "Python", "Java"},
		[]int{1, 2, 3, 4, 5},
		fmt.Errorf("发生错误"),
		User{Name: "张三", Age: 25},
	}

	for _, v := range values {
		fmt.Printf("%T => %s\n", v, format(v))
	}
}
