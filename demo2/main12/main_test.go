package main12

import (
	"fmt"
	"testing"
	"time"
)

func TestMain1(t *testing.T) {
	day := 3

	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("无效的日期")
	}
}

func TestMain2(t *testing.T) {
	day := 6

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("无效的日期")
	}
}

func TestMain3(t *testing.T) {
	score := 85

	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

func TestMain4(t *testing.T) {
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天是工作日")
	}
}

func TestMain5(t *testing.T) {
	num := 1

	switch num {
	case 1:
		fmt.Println("一")
		fallthrough
	case 2:
		fmt.Println("二")
		fallthrough
	case 3:
		fmt.Println("三")
	}
}

func TestMain6(t *testing.T) {
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)
}

func TestMain7(t *testing.T) {
	a := 10
	b := 3
	operator := "+"

	var result int
	valid := true

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("错误：除数不能为0")
			valid = false
		}
	case "%":
		if b != 0 {
			result = a % b
		} else {
			fmt.Println("错误：除数不能为0")
			valid = false
		}
	default:
		fmt.Println("错误：不支持的运算符")
		valid = false
	}

	if valid {
		fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
	}
}
