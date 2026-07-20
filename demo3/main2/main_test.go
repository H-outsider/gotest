package main2

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	// 使用make创建空map
	scores := make(map[string]int)
	scores["张三"] = 85
	scores["李四"] = 92
	fmt.Println("make创建:", scores)

	// 字面量初始化
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 28,
	}
	fmt.Println("字面量:", ages)

	// 空map
	empty := map[string]int{}
	fmt.Println("空map:", empty, "长度:", len(empty))
}

func TestMain2(t *testing.T) {
	m := make(map[string]int)

	// 写入
	m["apple"] = 5
	m["banana"] = 3
	m["orange"] = 8

	// 读取
	fmt.Println("apple:", m["apple"])
	fmt.Println("banana:", m["banana"])

	// 读取不存在的键，返回值类型的零值
	fmt.Println("grape:", m["grape"]) // 0

	// 获取长度
	fmt.Println("长度:", len(m))
}

func TestMain3(t *testing.T) {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	// 检查键是否存在
	value, exists := m["apple"]
	if exists {
		fmt.Println("apple存在，值为:", value)
	}

	value, exists = m["grape"]
	if !exists {
		fmt.Println("grape不存在")
	}

	// 简化写法
	if v, ok := m["banana"]; ok {
		fmt.Println("banana的值:", v)
	}
}

func TestMain4(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("删除前:", m)

	// 删除单个键
	delete(m, "b")
	fmt.Println("删除b后:", m)

	// 删除不存在的键不会报错
	delete(m, "x")
	fmt.Println("删除x后:", m)

	// 清空所有键值对（Go 1.21+）
	clear(m)
	fmt.Println("清空后:", m)
}

func TestMain5(t *testing.T) {
	fruits := map[string]int{
		"苹果": 5,
		"香蕉": 3,
		"橙子": 8,
	}

	// 遍历键值对
	fmt.Println("遍历键值对:")
	for key, value := range fruits {
		fmt.Printf("  %s: %d个\n", key, value)
	}

	// 只遍历键
	fmt.Println("\n只遍历键:")
	for key := range fruits {
		fmt.Println(" ", key)
	}
}

func TestMain6(t *testing.T) {
	text := "go is great go is simple go is fast"
	words := strings.Split(text, " ")

	// 统计词频
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// 输出结果
	fmt.Println("单词统计:")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d次\n", word, count)
	}

	// 查找出现最多的单词
	maxWord := ""
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			maxWord = word
			maxCount = count
		}
	}
	fmt.Printf("\n出现最多: %s (%d次)\n", maxWord, maxCount)
}
