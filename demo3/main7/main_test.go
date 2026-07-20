package main7

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	quotient, remainder := divmod(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)
}

func TestMain2(t *testing.T) {
	// 正常情况
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// 错误情况
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	}
}

func TestMain3(t *testing.T) {
	// 只要名字
	name, _, _ := getInfo()
	fmt.Println("姓名:", name)

	// 只要年龄
	_, age, _ := getInfo()
	fmt.Println("年龄:", age)

	// 全部获取
	name, age, email := getInfo()
	fmt.Printf("姓名:%s, 年龄:%d, 邮箱:%s\n", name, age, email)
}

func TestMain4(t *testing.T) {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max := minMax(nums)
	fmt.Printf("最小值: %d, 最大值: %d\n", min, max)
}

func TestMain5(t *testing.T) {
	if name, ok := findUser(2); ok {
		fmt.Println("找到用户:", name)
	} else {
		fmt.Println("用户不存在")
	}

	if name, ok := findUser(99); ok {
		fmt.Println("找到用户:", name)
	} else {
		fmt.Println("用户不存在")
	}
}

func TestMain6(t *testing.T) {
	testCases := []string{
		"localhost:8080",
		"192.168.1.1:3306",
		"invalid",
		"host:abc",
		"host:99999",
	}

	for _, addr := range testCases {
		host, port, err := parseAddress(addr)
		if err != nil {
			fmt.Printf("解析 %q 失败: %v\n", addr, err)
		} else {
			fmt.Printf("解析 %q 成功: host=%s, port=%d\n", addr, host, port)
		}
	}
}
