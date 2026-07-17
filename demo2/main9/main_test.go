package main9

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	age := 25
	hasLicense := true

	// 年龄>=18 且 有驾照 才能开车
	canDrive := age >= 18 && hasLicense
	fmt.Println("能否开车:", canDrive)

	// 真值表演示
	fmt.Println("\n=== 逻辑与真值表 ===")
	fmt.Println("true && true =", true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", false)
}

func TestMain2(t *testing.T) {
	isWeekend := true
	isHoliday := false

	// 周末或节假日可以休息
	canRest := isWeekend || isHoliday
	fmt.Println("能否休息:", canRest)

	// 真值表演示
	fmt.Println("\n=== 逻辑或真值表 ===")
	fmt.Println("true || true =", true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || true =", false || true)
	fmt.Println("false || false =", false)
}

func TestMain3(t *testing.T) {
	isRaining := false

	// 不下雨就可以出门
	canGoOut := !isRaining
	fmt.Println("能否出门:", canGoOut)

	// 取反演示
	fmt.Println("\n=== 逻辑非 ===")
	fmt.Println("!true =", !true)
	fmt.Println("!false =", !false)
}

func TestMain4(t *testing.T) {
	age := 20
	hasID := true
	isMember := false

	// 场景：进入酒吧需要年满18岁且带身份证，或者是会员
	canEnter := (age >= 18 && hasID) || isMember
	fmt.Println("能否进入:", canEnter)

	// 场景：不是VIP且余额不足100，显示充值提示
	isVip := false
	balance := 50.0
	showRecharge := !isVip && balance < 100
	fmt.Println("显示充值提示:", showRecharge)
}

func TestMain5(t *testing.T) {
	fmt.Println("=== 逻辑与短路 ===")
	// checkA()返回false，checkB()不会执行
	result1 := checkA() && checkB()
	fmt.Println("结果:", result1)

	fmt.Println("\n=== 逻辑或短路 ===")
	// 交换顺序，checkB()返回true，checkA()不会执行
	result2 := checkB() || checkA()
	fmt.Println("结果:", result2)
}

func TestMain6(t *testing.T) {
	// 用户输入
	username := "admin"
	password := "123456"
	isAccountLocked := false
	loginAttempts := 2

	// 验证逻辑
	isUsernameValid := len(username) >= 3
	isPasswordValid := len(password) >= 6
	isAccountActive := !isAccountLocked
	hasAttemptsLeft := loginAttempts < 5

	// 能登录的条件：用户名有效 且 密码有效 且 账户未锁定 且 还有尝试次数
	canLogin := isUsernameValid && isPasswordValid && isAccountActive && hasAttemptsLeft

	fmt.Println("===== 登录验证 =====")
	fmt.Println("用户名有效:", isUsernameValid)
	fmt.Println("密码有效:", isPasswordValid)
	fmt.Println("账户未锁定:", isAccountActive)
	fmt.Println("还有尝试次数:", hasAttemptsLeft)
	fmt.Println("--------------------")
	fmt.Println("能否登录:", canLogin)

	// 显示错误提示的逻辑
	if !canLogin {
		if !isUsernameValid || !isPasswordValid {
			fmt.Println("提示: 用户名或密码格式不正确")
		} else if !isAccountActive {
			fmt.Println("提示: 账户已被锁定")
		} else if !hasAttemptsLeft {
			fmt.Println("提示: 尝试次数过多，请稍后再试")
		}
	}
}

func TestMain7(t *testing.T) {
	a, b, c := true, false, true

	// ! 优先级最高
	result1 := !a || b
	fmt.Println("!true || false =", result1) // false || false = false

	// && 优先级高于 ||
	result2 := a || b && c
	fmt.Println("true || false && true =", result2) // true || false = true
	// 等价于 true || (false && true)

	// 使用括号改变优先级
	result3 := (a || b) && c
	fmt.Println("(true || false) && true =", result3) // true && true = true
}
