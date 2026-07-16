package main2

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	name := "张三"
	age := 25
	score := 98.5
	isPass := true

	fmt.Println(name, age, score, isPass)
}

func TestMain2(t *testing.T) {
	a, b, c := 1, 2, 3
	name, age := "李四", 30
	x, y, z := 1, "hello", true

	fmt.Println(a, b, c)
	fmt.Println(name, age)
	fmt.Println(x, y, z)
}

func TestMain3(t *testing.T) {
	if num := 9; num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "是一位数")
	} else {
		fmt.Println(num, "是多位数")
	}

	// 注意：num在if块外面不可用
	// fmt.Println(num) // 这行会报错
}

func TestMain4(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// i在循环外不可用
}

func TestMain5(t *testing.T) {
	a := 1
	fmt.Println("a =", a)

	// b是新变量，a已存在，这是合法的
	a, b := 2, 3
	fmt.Println("a =", a, ", b =", b)

	// c是新变量，a和b已存在，这也是合法的
	a, b, c := 4, 5, 6
	fmt.Println("a =", a, ", b =", b, ", c =", c)
}

func TestMain6(t *testing.T) {
	// 商品信息
	productName := "Go语言编程"
	price := 79.00
	quantity := 2

	// 计算
	subtotal := price * float64(quantity)
	discount := 0.9 // 9折
	finalPrice := subtotal * discount

	// 输出订单信息
	fmt.Println("===== 订单详情 =====")
	fmt.Println("商品:", productName)
	fmt.Printf("单价: %.2f元\n", price)
	fmt.Println("数量:", quantity)
	fmt.Printf("小计: %.2f元\n", subtotal)
	fmt.Printf("折扣: %.0f%%\n", discount*100)
	fmt.Printf("实付: %.2f元\n", finalPrice)
}
