package main6

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func greet(name string) {
	fmt.Println("你好,", name)
}

func printSum(a, b int) {
	fmt.Println("和:", a+b)
}

func addInt(a int, b int) int {
	return a + b
}

func multiplyInt(a, b int) int {
	return a * b
}

func divideInt(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func rectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func addForOperation(a, b int) int {
	return a + b
}

// 加法
func addFloat(a, b float64) float64 {
	return a + b
}

// 减法
func subtractFloat(a, b float64) float64 {
	return a - b
}

// 乘法
func multiplyFloat(a, b float64) float64 {
	return a * b
}

// 除法，返回结果和错误信息
func divideFloat(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// calculate 计算器函数。
func calculate(a, b float64, operator string) {
	var result float64
	var err error

	switch operator {
	case "+":
		result = addFloat(a, b)
	case "-":
		result = subtractFloat(a, b)
	case "*":
		result = multiplyFloat(a, b)
	case "/":
		result, err = divideFloat(a, b)
		if err != nil {
			fmt.Println("错误:", err)
			return
		}
	default:
		fmt.Println("不支持的操作符:", operator)
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, operator, b, result)
}

func modifyValue(x int) {
	x = 100
}

func modifyPointer(x *int) {
	*x = 100
}
