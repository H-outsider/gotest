package main4

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var s BasicSpeaker

	// 初始状态：type和value都是nil
	fmt.Printf("初始: type=%T, value=%v\n", s, s)

	// 赋值Dog
	s = BasicDog{Name: "Buddy"}
	fmt.Printf("Dog: type=%T, value=%v\n", s, s)
	s.Speak()

	// 赋值Cat
	s = BasicCat{Name: "Kitty"}
	fmt.Printf("Cat: type=%T, value=%v\n", s, s)
	s.Speak()
}

func TestMain2(t *testing.T) {
	var s NilSpeaker

	// 情况1：接口本身是nil
	fmt.Println("s == nil:", s == nil) // true
	// s.Speak()  // 会panic！

	// 情况2：接口不是nil，但动态值是nil
	var d *NilDog
	s = d
	fmt.Println("赋值nil指针后:")
	fmt.Println("s == nil:", s == nil) // false！
	fmt.Printf("type=%T, value=%v\n", s, s)
	s.Speak() // 不会panic，会调用方法
}

func TestMain3(t *testing.T) {
	var s1, s2 ComparableSpeaker

	// 两个nil接口相等
	fmt.Println("两个nil接口:", s1 == s2) // true

	// 相同类型相同值
	s1 = ComparableDog{Name: "Buddy"}
	s2 = ComparableDog{Name: "Buddy"}
	fmt.Println("相同Dog:", s1 == s2) // true

	// 相同类型不同值
	s2 = ComparableDog{Name: "Max"}
	fmt.Println("不同Dog:", s1 == s2) // false

	// 不同类型
	type Cat struct{ Name string }
	// s2 = Cat{Name: "Buddy"}  // Cat没有Speak方法
}

func TestMain4(t *testing.T) {
	var c1, c2 Container

	c1 = SliceContainer{Data: []int{1, 2, 3}}
	c2 = SliceContainer{Data: []int{1, 2, 3}}

	// 切片不可比较，比较接口值会panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("比较panic:", r)
		}
	}()

	fmt.Println(c1 == c2) // panic!
}

func TestMain5(t *testing.T) {
	var a Animal

	fmt.Println("=== nil接口 ===")
	inspectInterface(a)

	fmt.Println("=== Dog值 ===")
	a = InspectDog{Name: "Buddy", Age: 3}
	inspectInterface(a)

	fmt.Println("=== Cat指针 ===")
	a = &InspectCat{Name: "Kitty"}
	inspectInterface(a)

	fmt.Println("=== nil Cat指针 ===")
	var c *InspectCat
	a = c
	inspectInterface(a)
}
