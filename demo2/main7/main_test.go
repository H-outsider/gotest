package main7

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMain1(t *testing.T) {
	var i int = 100
	var i64 int64 = int64(i)   // int -> int64
	var f float64 = float64(i) // int -> float64
	var u uint = uint(i)       // int -> uint

	fmt.Printf("int: %d\n", i)
	fmt.Printf("int64: %d\n", i64)
	fmt.Printf("float64: %f\n", f)
	fmt.Printf("uint: %d\n", u)
}

func TestMain2(t *testing.T) {
	var f1 float64 = 3.7
	var f2 float64 = -3.7

	i1 := int(f1) // 截断小数部分
	i2 := int(f2)

	fmt.Printf("3.7 -> int: %d\n", i1)
	fmt.Printf("-3.7 -> int: %d\n", i2)
}

func TestMain3(t *testing.T) {
	s := "Hello, 世界"

	// 字符串转字节切片
	bytes := []byte(s)
	fmt.Println("字节切片:", bytes)

	// 字节切片转字符串
	s2 := string(bytes)
	fmt.Println("字符串:", s2)

	// 字符串转rune切片
	runes := []rune(s)
	fmt.Println("rune切片:", runes)

	// rune切片转字符串
	s3 := string(runes)
	fmt.Println("字符串:", s3)
}

func TestMain4(t *testing.T) {
	// 整数转字符串
	num := 42
	s1 := strconv.Itoa(num) // Itoa = Integer to ASCII
	fmt.Println("整数转字符串:", s1)

	// 字符串转整数
	s2 := "123"
	num2, err := strconv.Atoi(s2) // Atoi = ASCII to Integer
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("字符串转整数:", num2)
	}

	// 更通用的方式
	s3 := strconv.FormatInt(int64(num), 10) // 第二个参数是进制
	fmt.Println("FormatInt:", s3)

	num3, _ := strconv.ParseInt("456", 10, 64) // 字符串、进制、位数
	fmt.Println("ParseInt:", num3)
}

func TestMAin5(t *testing.T) {
	// 浮点数转字符串
	f := 3.14159
	s1 := strconv.FormatFloat(f, 'f', 2, 64) // 格式、精度、位数
	fmt.Println("浮点转字符串:", s1)

	// 字符串转浮点数
	s2 := "2.71828"
	f2, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("字符串转浮点:", f2)
	}

	// 使用Sprintf也可以
	s3 := fmt.Sprintf("%.2f", f)
	fmt.Println("Sprintf:", s3)
}

func TestMain6(t *testing.T) {
	var i int = 100
	var i64 int64 = 100

	// 下面这行在其他语言可能可以，但Go会报错
	// result := i + i64  // 编译错误：invalid operation

	// 必须显式转换
	result := int64(i) + i64
	fmt.Println(result)
}

func TestMain7(t *testing.T) {
	// 模拟从配置文件读取的字符串数据
	configPort := "8080"
	configTimeout := "30.5"
	configDebug := "true"

	// 转换为实际类型
	port, err := strconv.Atoi(configPort)
	if err != nil {
		fmt.Println("端口转换失败")
		return
	}

	timeout, err := strconv.ParseFloat(configTimeout, 64)
	if err != nil {
		fmt.Println("超时时间转换失败")
		return
	}

	debug, err := strconv.ParseBool(configDebug)
	if err != nil {
		fmt.Println("调试标志转换失败")
		return
	}

	fmt.Println("===== 配置信息 =====")
	fmt.Printf("端口: %d (类型: %T)\n", port, port)
	fmt.Printf("超时: %.1f秒 (类型: %T)\n", timeout, timeout)
	fmt.Printf("调试模式: %t (类型: %T)\n", debug, debug)
}
