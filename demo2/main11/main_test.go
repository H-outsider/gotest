package main11

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	age := 15

	if age >= 18 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}
}

func TestMain2(t *testing.T) {
	score := 85

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}

func TestMain3(t *testing.T) {
	month := 8

	if month >= 3 && month <= 5 {
		fmt.Println("春季")
	} else if month >= 6 && month <= 8 {
		fmt.Println("夏季")
	} else if month >= 9 && month <= 11 {
		fmt.Println("秋季")
	} else if month == 12 || month == 1 || month == 2 {
		fmt.Println("冬季")
	} else {
		fmt.Println("无效月份")
	}
}

func TestMain4(t *testing.T) {
	if num := 9; num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "是一位数")
	} else {
		fmt.Println(num, "是多位数")
	}
}

func TestMain5(t *testing.T) {
	height := 1.75 // 米
	weight := 70.0 // 公斤

	// 计算BMI
	bmi := weight / (height * height)

	fmt.Printf("身高: %.2f米\n", height)
	fmt.Printf("体重: %.1f公斤\n", weight)
	fmt.Printf("BMI: %.1f\n", bmi)
	fmt.Print("评估: ")

	if bmi < 18.5 {
		fmt.Println("偏瘦，建议适当增加营养")
	} else if bmi < 24 {
		fmt.Println("正常，请继续保持")
	} else if bmi < 28 {
		fmt.Println("偏胖，建议适当运动")
	} else {
		fmt.Println("肥胖，建议咨询医生")
	}
}

func TestMain6(t *testing.T) {
	score := 95

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}

func TestMain7(t *testing.T) {
	age := 20
	var status string

	// Go没有三元运算符，必须这样写
	if age >= 18 {
		status = "成年"
	} else {
		status = "未成年"
	}

	fmt.Println(status)
}
