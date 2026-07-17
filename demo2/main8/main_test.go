package main8

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	a := 10
	b := 3

	fmt.Println("a =", a, ", b =", b)
	fmt.Println("a + b =", a+b) // 加法
	fmt.Println("a - b =", a-b) // 减法
	fmt.Println("a * b =", a*b) // 乘法
	fmt.Println("a / b =", a/b) // 整数除法，结果为3
	fmt.Println("a % b =", a%b) // 取模，余数为1
}

func TestMain2(t *testing.T) {
	// 整数除法：小数部分被截断
	a := 7
	b := 2
	fmt.Println("整数除法 7/2 =", a/b) // 结果是3，不是3.5

	// 浮点除法：保留小数
	x := 7.0
	y := 2.0
	fmt.Println("浮点除法 7.0/2.0 =", x/y) // 结果是3.5

	// 混合类型需要转换
	result := float64(a) / float64(b)
	fmt.Println("转换后 7/2 =", result)
}

func TestMain3(t *testing.T) {
	// 判断奇偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "是偶数")
		} else {
			fmt.Println(i, "是奇数")
		}
	}

	// 循环索引（如轮播图）
	total := 5 // 总共5张图
	for i := 0; i < 10; i++ {
		index := i % total
		fmt.Printf("第%d次显示第%d张图\n", i+1, index+1)
	}
}

func TestMain4(t *testing.T) {
	a := 10
	b := -a // 负号
	c := +a // 正号（通常省略）

	fmt.Println("a =", a)
	fmt.Println("-a =", b)
	fmt.Println("+a =", c)
}

func TestMain5(t *testing.T) {
	a := 10
	fmt.Println("初始值: a =", a)

	a += 5 // a = a + 5
	fmt.Println("a += 5 后: a =", a)

	a -= 3 // a = a - 3
	fmt.Println("a -= 3 后: a =", a)

	a *= 2 // a = a * 2
	fmt.Println("a *= 2 后: a =", a)

	a /= 4 // a = a / 4
	fmt.Println("a /= 4 后: a =", a)

	a %= 3 // a = a % 3
	fmt.Println("a %= 3 后: a =", a)
}

func TestMain6(t *testing.T) {
	a := 10

	a++ // a = a + 1
	fmt.Println("a++ 后:", a)

	a-- // a = a - 1
	fmt.Println("a-- 后:", a)

	// 注意：以下写法在Go中是错误的
	// b := a++     // 错误：a++是语句，不能作为表达式
	// fmt.Println(++a)  // 错误：Go没有前置++
}

func TestMain7(t *testing.T) {
	// 乘法优先级高于加法
	result1 := 2 + 3*4
	fmt.Println("2 + 3*4 =", result1) // 14，不是20

	// 使用括号改变优先级
	result2 := (2 + 3) * 4
	fmt.Println("(2 + 3) * 4 =", result2) // 20

	// 同优先级从左到右
	result3 := 20 / 4 * 2
	fmt.Println("20 / 4 * 2 =", result3) // 10，从左到右计算
}

func TestMain8(t *testing.T) {
	// 商品价格
	bookPrice := 59.9
	penPrice := 9.9
	notebookPrice := 15.0

	// 购买数量
	bookQty := 2
	penQty := 5
	notebookQty := 3

	// 计算各项小计
	bookTotal := bookPrice * float64(bookQty)
	penTotal := penPrice * float64(penQty)
	notebookTotal := notebookPrice * float64(notebookQty)

	// 计算总价
	subtotal := bookTotal + penTotal + notebookTotal

	// 满100减20
	discount := 0.0
	if subtotal >= 100 {
		discount = 20.0
	}

	finalPrice := subtotal - discount

	// 输出账单
	fmt.Println("===== 购物清单 =====")
	fmt.Printf("书籍: %.2f x %d = %.2f\n", bookPrice, bookQty, bookTotal)
	fmt.Printf("钢笔: %.2f x %d = %.2f\n", penPrice, penQty, penTotal)
	fmt.Printf("笔记本: %.2f x %d = %.2f\n", notebookPrice, notebookQty, notebookTotal)
	fmt.Println("--------------------")
	fmt.Printf("小计: %.2f\n", subtotal)
	fmt.Printf("优惠: -%.2f\n", discount)
	fmt.Printf("实付: %.2f\n", finalPrice)
}
