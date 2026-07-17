package main14

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}

	// 获取索引和值
	fmt.Println("索引和值:")
	for i, v := range nums {
		fmt.Printf("索引%d: 值%d\n", i, v)
	}

	// 只获取值
	fmt.Println("\n只要值:")
	for _, v := range nums {
		fmt.Println(v)
	}

	// 只获取索引
	fmt.Println("\n只要索引:")
	for i := range nums {
		fmt.Println(i)
	}
}

func TestMain2(t *testing.T) {
	s := "Go语言"

	fmt.Println("遍历字符串:")
	for i, r := range s {
		fmt.Printf("位置%d: 字符'%c' (码点%d)\n", i, r, r)
	}
}

func TestMain3(t *testing.T) {
	scores := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
	}

	fmt.Println("遍历map:")
	for name, score := range scores {
		fmt.Printf("%s: %d分\n", name, score)
	}

	fmt.Println("\n只遍历键:")
	for name := range scores {
		fmt.Println(name)
	}
}

func TestMain4(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	fmt.Println("元素:", numbers)
	fmt.Println("总和:", sum)
	fmt.Printf("平均值: %.1f\n", float64(sum)/float64(len(numbers)))
}

func TestMain5(t *testing.T) {
	text := "hello world"

	// 用map统计字符频率
	charCount := make(map[rune]int)

	for _, char := range text {
		charCount[char]++
	}

	fmt.Println("原文:", text)
	fmt.Println("\n字符统计:")
	for char, count := range charCount {
		if char == ' ' {
			fmt.Printf("空格: %d次\n", count)
		} else {
			fmt.Printf("'%c': %d次\n", char, count)
		}
	}
}

func TestMain6(t *testing.T) {
	fruits := []string{"苹果", "香蕉", "橙子"}

	// 普通for循环
	fmt.Println("普通for:")
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	// range循环
	fmt.Println("\nrange:")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
