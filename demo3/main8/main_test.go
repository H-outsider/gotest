package main8

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	fmt.Println(sum(1, 2, 3))       // 6
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
	fmt.Println(sum())              // 0，没有参数也可以
	fmt.Println(sum(10))            // 10，单个参数
}

func TestMain2(t *testing.T) {
	printInfo(1, 2, 3)
}

func TestMain3(t *testing.T) {
	fmt.Println(join("-", "a", "b", "c"))
	fmt.Println(join(", ", "苹果", "香蕉", "橙子"))
	fmt.Println(join(":", "usr", "local", "bin"))
}

func TestMain4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	// 使用...展开切片
	result := sum(nums...)
	fmt.Println("切片求和:", result)

	// 也可以直接传递值
	result2 := sum(10, 20, 30)
	fmt.Println("直接传值:", result2)
}

func TestMain5(t *testing.T) {
	printAll("hello", 42, 3.14, true, []int{1, 2, 3})
}

func TestMain6(t *testing.T) {
	log(INFO, "服务启动")
	log(INFO, "监听端口: %d", 8080)
	log(WARN, "内存使用率: %.1f%%", 85.5)
	log(ERROR, "连接失败: %s, 重试次数: %d", "数据库", 3)
}
