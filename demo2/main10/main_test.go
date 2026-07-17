package main10

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	age := 20

	if age >= 18 {
		fmt.Println("你已成年")
	}

	fmt.Println("程序继续执行")
}

func TestMain2(t *testing.T) {
	// score在if语句中声明，只在if块内有效
	if score := 85; score >= 60 {
		fmt.Println("及格了，分数:", score)
	}

	// fmt.Println(score)  // 错误：score在这里不可用
}

func TestMain3(t *testing.T) {
	age := 25
	hasID := true

	// 使用 && (与)
	if age >= 18 && hasID {
		fmt.Println("可以进入")
	}

	isVip := true
	isHoliday := false

	// 使用 || (或)
	if isVip || isHoliday {
		fmt.Println("享受折扣")
	}

	isBlocked := false

	// 使用 ! (非)
	if !isBlocked {
		fmt.Println("账户正常")
	}
}

func TestMain4(t *testing.T) {
	age := 25
	hasLicense := true

	if age >= 18 {
		fmt.Println("年龄符合要求")

		if hasLicense {
			fmt.Println("可以开车")
		}
	}
}

func TestMain5(t *testing.T) {
	// 模拟用户输入
	inputUsername := "admin"
	inputPassword := "123456"

	// 系统中存储的正确信息
	correctUsername := "admin"
	correctPassword := "123456"
	isAccountActive := true

	// 验证用户名和密码
	if inputUsername == correctUsername && inputPassword == correctPassword {
		// 检查账户状态
		if isAccountActive {
			fmt.Println("登录成功！欢迎回来，", inputUsername)
		}
	}
}

func TestMain6(t *testing.T) {
	a, b := 10, 20

	if a == b {
		fmt.Println("a等于b")
	}

	if a != b {
		fmt.Println("a不等于b")
	}

	if a < b {
		fmt.Println("a小于b")
	}

	if a <= b {
		fmt.Println("a小于等于b")
	}
}
