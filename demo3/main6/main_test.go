package main6

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	sayHello()
	sayHello()
}

func TestMain2(t *testing.T) {
	greet("张三")
	greet("李四")
	printSum(10, 20)
}

func TestMain3(t *testing.T) {
	sum := addInt(3, 5)
	fmt.Println("3 + 5 =", sum)

	product := multiplyInt(4, 6)
	fmt.Println("4 * 6 =", product)

	// 直接使用返回值
	fmt.Println("10 + 20 =", addInt(10, 20))
}

func TestMain4(t *testing.T) {
	q, r := divideInt(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", q, r)

	// 忽略某个返回值
	q2, _ := divideInt(20, 3)
	fmt.Println("20 ÷ 3 =", q2)
}

func TestMain5(t *testing.T) {
	a, p := rectangle(5, 3)
	fmt.Println("面积:", a)
	fmt.Println("周长:", p)
}

func TestMain6(t *testing.T) {
	// 函数赋值给变量
	operation := addForOperation
	result := operation(3, 4)
	fmt.Println("结果:", result)

	// 匿名函数
	subtract := func(a, b int) int {
		return a - b
	}
	fmt.Println("10 - 3 =", subtract(10, 3))
}

func TestMain7(t *testing.T) {
	calculate(10, 3, "+")
	calculate(10, 3, "-")
	calculate(10, 3, "*")
	calculate(10, 3, "/")
	calculate(10, 0, "/")
}

func TestMain8(t *testing.T) {
	a := 10
	modifyValue(a)
	fmt.Println("值传递后:", a) // 10

	modifyPointer(&a)
	fmt.Println("指针传递后:", a) // 100
}
