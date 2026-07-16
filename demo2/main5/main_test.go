package main5

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var price float64 = 99.99
	var rate float32 = 0.85

	fmt.Println("价格:", price)
	fmt.Println("折扣率:", rate)
}

func TestMain2(t *testing.T) {
	var big float64 = 1.23e10   // 1.23 × 10^10
	var small float64 = 4.56e-8 // 4.56 × 10^-8

	fmt.Printf("大数: %f\n", big)
	fmt.Printf("大数(科学计数法): %e\n", big)
	fmt.Printf("小数: %f\n", small)
	fmt.Printf("小数(科学计数法): %e\n", small)
}

func TestMain3(t *testing.T) {
	pi := 3.141592653589793

	fmt.Printf("默认: %f\n", pi)     // 默认6位小数
	fmt.Printf("2位小数: %.2f\n", pi) // 保留2位
	fmt.Printf("8位小数: %.8f\n", pi) // 保留8位
	fmt.Printf("科学计数法: %e\n", pi)
	fmt.Printf("自动选择: %g\n", pi) // 根据情况选择%f或%e
}

func TestMain4(t *testing.T) {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = a + b

	fmt.Println("0.1 + 0.2 =", c)
	fmt.Println("等于0.3吗?", c == 0.3)

	fmt.Printf("精确值: %.20f\n", c)
}

func TestMain5(t *testing.T) {
	isLogin := true
	isVip := false

	fmt.Println("是否登录:", isLogin)
	fmt.Println("是否VIP:", isVip)

	// 布尔值可以直接用于条件判断
	if isLogin {
		fmt.Println("欢迎回来！")
	}

	if !isVip {
		fmt.Println("升级VIP享受更多特权")
	}
}

func TestMain6(t *testing.T) {
	a := 10
	b := 20

	fmt.Println("a == b:", a == b) // 等于
	fmt.Println("a != b:", a != b) // 不等于
	fmt.Println("a < b:", a < b)   // 小于
	fmt.Println("a > b:", a > b)   // 大于
	fmt.Println("a <= b:", a <= b) // 小于等于
	fmt.Println("a >= b:", a >= b) // 大于等于
}

func TestMain7(t *testing.T) {
	// 商品信息
	productName := "Go语言实战"
	originalPrice := 89.00
	discountRate := 0.8
	isOnSale := true

	// 计算最终价格
	var finalPrice float64
	if isOnSale {
		finalPrice = originalPrice * discountRate
	} else {
		finalPrice = originalPrice
	}

	// 判断是否包邮
	freeShipping := finalPrice >= 50.0

	// 输出信息
	fmt.Println("===== 订单信息 =====")
	fmt.Println("商品名称:", productName)
	fmt.Printf("原价: %.2f元\n", originalPrice)
	fmt.Println("是否促销:", isOnSale)
	if isOnSale {
		fmt.Printf("折扣: %.0f%%\n", discountRate*100)
	}
	fmt.Printf("实付: %.2f元\n", finalPrice)
	fmt.Println("是否包邮:", freeShipping)
}
