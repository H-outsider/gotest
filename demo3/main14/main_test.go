package main14

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMain1(t *testing.T) {
	// 字符串转int
	s := "42"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("Atoi结果:", n, "类型:", fmt.Sprintf("%T", n))
	}

	// int转字符串
	num := 123
	str := strconv.Itoa(num)
	fmt.Println("Itoa结果:", str)

	// 转换失败的情况
	bad := "abc"
	_, err = strconv.Atoi(bad)
	if err != nil {
		fmt.Println("无效数字:", err)
	}
}

func TestMain2(t *testing.T) {
	// ParseInt(s string, base int, bitSize int)
	// base: 进制（0表示自动推断）
	// bitSize: 结果的位数（0, 8, 16, 32, 64）

	// 十进制
	n1, _ := strconv.ParseInt("42", 10, 64)
	fmt.Println("十进制:", n1)

	// 二进制
	n2, _ := strconv.ParseInt("101010", 2, 64)
	fmt.Println("二进制101010:", n2)

	// 十六进制
	n3, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println("十六进制ff:", n3)

	// 八进制
	n4, _ := strconv.ParseInt("77", 8, 64)
	fmt.Println("八进制77:", n4)

	// 自动推断进制（0x前缀为16进制，0前缀为8进制）
	n5, _ := strconv.ParseInt("0xff", 0, 64)
	fmt.Println("自动推断0xff:", n5)

	// 负数
	n6, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println("负数:", n6)
}

func TestMain3(t *testing.T) {
	// 字符串转float64
	s := "3.14159"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("ParseFloat结果:", f)
	}

	// 科学计数法
	s2 := "1.5e10"
	f2, _ := strconv.ParseFloat(s2, 64)
	fmt.Println("科学计数法:", f2)

	// float64转字符串
	num := 3.14159265358979

	// 'f' 普通格式，2位小数
	str1 := strconv.FormatFloat(num, 'f', 2, 64)
	fmt.Println("2位小数:", str1)

	// 'f' 普通格式，-1表示最少位数
	str2 := strconv.FormatFloat(num, 'f', -1, 64)
	fmt.Println("最少位数:", str2)

	// 'e' 科学计数法
	str3 := strconv.FormatFloat(num, 'e', 3, 64)
	fmt.Println("科学计数法:", str3)

	// 'g' 自动选择
	str4 := strconv.FormatFloat(num, 'g', 5, 64)
	fmt.Println("自动选择:", str4)
}

func TestMain4(t *testing.T) {
	// 字符串转bool
	// 接受: "1", "t", "T", "true", "TRUE", "True"
	//       "0", "f", "F", "false", "FALSE", "False"
	values := []string{"true", "false", "1", "0", "T", "F", "TRUE", "FALSE"}

	for _, v := range values {
		b, err := strconv.ParseBool(v)
		if err != nil {
			fmt.Printf("%q -> 错误: %v\n", v, err)
		} else {
			fmt.Printf("%q -> %v\n", v, b)
		}
	}

	// bool转字符串
	fmt.Println("\nbool转字符串:")
	fmt.Println("true ->", strconv.FormatBool(true))
	fmt.Println("false ->", strconv.FormatBool(false))
}

func TestMain5(t *testing.T) {
	// 添加引号
	s := "hello\tworld"
	quoted := strconv.Quote(s)
	fmt.Println("Quote:", quoted)

	// 去除引号
	unquoted, _ := strconv.Unquote(quoted)
	fmt.Println("Unquote:", unquoted)

	// 检查是否可打印
	fmt.Println("IsPrint('a'):", strconv.IsPrint('a'))
	fmt.Println("IsPrint('\\t'):", strconv.IsPrint('\t'))
}

func TestMain6(t *testing.T) {
	config := Config{
		Port:    8080,
		Debug:   false,
		Timeout: 30.0,
	}

	// 模拟参数
	args := map[string]string{
		"port":    "3000",
		"debug":   "true",
		"timeout": "60.5",
	}

	for key, value := range args {
		if err := parseArg(key, value, &config); err != nil {
			fmt.Println("错误:", err)
		}
	}

	fmt.Printf("配置: Port=%d, Debug=%v, Timeout=%.1f\n",
		config.Port, config.Debug, config.Timeout)
}

func TestMain7(t *testing.T) {
	buf := make([]byte, 0, 64)

	buf = append(buf, "数字: "...)
	buf = strconv.AppendInt(buf, 42, 10)
	buf = append(buf, ", 浮点: "...)
	buf = strconv.AppendFloat(buf, 3.14, 'f', 2, 64)
	buf = append(buf, ", 布尔: "...)
	buf = strconv.AppendBool(buf, true)

	fmt.Println(string(buf))
}
