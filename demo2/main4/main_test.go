package main4

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	var a int = 100
	var b int8 = 127
	var c int64 = 9223372036854775807
	var d uint = 200

	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int64:", c)
	fmt.Println("uint:", d)
}

func TestMain2(t *testing.T) {
	decimal := 42      // 十进制
	binary := 0b101010 // 二进制，0b或0B开头
	octal := 0o52      // 八进制，0o或0O开头
	hex := 0x2A        // 十六进制，0x或0X开头

	fmt.Println("十进制 42:", decimal)
	fmt.Println("二进制 0b101010:", binary)
	fmt.Println("八进制 0o52:", octal)
	fmt.Println("十六进制 0x2A:", hex)

	// 大数字可以用下划线分隔，增加可读性
	bigNumber := 1_000_000_000
	fmt.Println("十亿:", bigNumber)
}

func TestMain3(t *testing.T) {
	// byte是uint8的别名，通常用于处理ASCII字符或二进制数据
	var b byte = 'A'
	fmt.Printf("byte: %c, 数值: %d\n", b, b)

	// rune是int32的别名，用于表示Unicode字符
	var r rune = '中'
	fmt.Printf("rune: %c, Unicode码点: %d\n", r, r)
}

func TestMain4(t *testing.T) {
	var signed int8 = -100   // 有符号，可以是负数
	var unsigned uint8 = 200 // 无符号，只能是非负数

	fmt.Println("有符号:", signed)
	fmt.Println("无符号:", unsigned)

	// 无符号整数不能存储负数
	// var wrong uint8 = -1  // 编译错误
}

func TestMain5(t *testing.T) {
	var a int8 = 127
	fmt.Println("int8最大值:", a)

	a = a + 1              // 溢出
	fmt.Println("加1后:", a) // 变成-128

	var b uint8 = 255
	fmt.Println("uint8最大值:", b)

	b = b + 1              // 溢出
	fmt.Println("加1后:", b) // 变成0
}

func TestMain6(t *testing.T) {
	// 用户信息
	var userId int64 = 10000000001 // 用户ID，可能很大
	var age int8 = 28              // 年龄，不会超过127
	var balance int64 = 1234567890 // 余额（分），需要大范围
	var loginCount uint32 = 1500   // 登录次数，非负数

	fmt.Println("===== 用户信息 =====")
	fmt.Printf("用户ID: %d\n", userId)
	fmt.Printf("年龄: %d岁\n", age)
	fmt.Printf("余额: %.2f元\n", float64(balance)/100)
	fmt.Printf("登录次数: %d次\n", loginCount)

	// 类型大小
	fmt.Println("\n===== 类型大小 =====")
	fmt.Printf("int8占用: 1字节\n")
	fmt.Printf("int32占用: 4字节\n")
	fmt.Printf("int64占用: 8字节\n")
}

func TestMain7(t *testing.T) {
	var a int = 100
	var b int64 = 100

	// a = b       // 编译错误，类型不同
	a = int(b) // 必须显式转换

	fmt.Println(a)
}
