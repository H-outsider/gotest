package main1

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var name string // 声明一个字符串变量
	name = "张三"     // 给变量赋值
	fmt.Println(name)
}

func TestMain2(t *testing.T) {
	var age int = 25
	fmt.Println(age)
}

func TestMain3(t *testing.T) {
	var name = "李四"     // 自动推断为string
	var age = 30        // 自动推断为int
	var price = 19.99   // 自动推断为float64
	var isActive = true // 自动推断为bool

	fmt.Println(name, age, price, isActive)
}
func TestMain4(t *testing.T) {
	// 同一行声明多个同类型变量
	var a, b, c int = 1, 2, 3

	// 使用括号批量声明不同类型变量
	var (
		name   string  = "王五"
		age    int     = 28
		salary float64 = 8500.50
	)

	fmt.Println(a, b, c)
	fmt.Println(name, age, salary)
}

func TestMain5(t *testing.T) {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("int的零值: %d\n", i)
	fmt.Printf("float64的零值: %f\n", f)
	fmt.Printf("bool的零值: %t\n", b)
	fmt.Printf("string的零值: %q\n", s)
}

func TestMain6(t *testing.T) {
	// 用户基本信息
	var (
		username string  = "zhangsan"
		nickname string  = "张三"
		age      int     = 25
		balance  float64 = 1000.50
		isVip    bool    = true
	)

	// 打印用户信息
	fmt.Println("===== 用户信息 =====")
	fmt.Println("用户名:", username)
	fmt.Println("昵称:", nickname)
	fmt.Println("年龄:", age)
	fmt.Printf("余额: %.2f元\n", balance)
	fmt.Println("VIP会员:", isVip)

	// 修改变量的值
	age = 26
	balance = balance + 500
	fmt.Println("\n===== 更新后 =====")
	fmt.Println("年龄:", age)
	fmt.Printf("余额: %.2f元\n", balance)
}
