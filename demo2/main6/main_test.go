package main6

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestMain1(t *testing.T) {
	s1 := "Hello, Go!"
	s2 := "第一行\n第二行"        // \n 换行
	s3 := "路径: C:\\Go\\bin" // \\ 反斜杠
	s4 := "他说:\"Go很棒\""     // \" 双引号

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func TestMain2(t *testing.T) {
	// 原始字符串，保留换行和格式
	sql := `SELECT * 
	FROM users 
	WHERE age > 18`

	// 保留反斜杠，不转义
	path := `C:\Go\bin\go.exe`

	fmt.Println(sql)
	fmt.Println(path)
}

func TestMain3(t *testing.T) {
	firstName := "张"
	lastName := "三"

	// 使用 + 拼接
	fullName := firstName + lastName
	fmt.Println("姓名:", fullName)

	// 使用 fmt.Sprintf 格式化
	age := 25
	info := fmt.Sprintf("%s, %d岁", fullName, age)
	fmt.Println("信息:", info)
}

func TestMain4(t *testing.T) {
	s := "Hello,你好"

	// len() 返回字节数
	fmt.Println("字节长度:", len(s))

	// utf8.RuneCountInString() 返回字符数
	fmt.Println("字符长度:", utf8.RuneCountInString(s))
}

func TestMain5(t *testing.T) {
	s := "Go语言"

	// 按字节遍历
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %x\n", i, s[i])
	}
}

func TestMain6(t *testing.T) {
	s := "Go语言"

	// 使用range按字符遍历
	for i, r := range s {
		fmt.Printf("索引%d: %c (Unicode: %d)\n", i, r, r)
	}
}

func TestMain7(t *testing.T) {
	s := "Hello"
	// s[0] = 'h'  // 编译错误：cannot assign to s[0]

	// 如果需要修改，要创建新字符串
	s = "hello" // 这是赋值新字符串，不是修改原字符串
	fmt.Println(s)

	// 或者转换为[]byte修改后再转回来
	bytes := []byte(s)
	bytes[0] = 'H'
	s = string(bytes)
	fmt.Println(s)
}

func TestMain8(t *testing.T) {
	// byte 处理单字节字符
	var b byte = 'A'
	fmt.Printf("byte: %c, 值: %d\n", b, b)

	// rune 处理任意Unicode字符
	var r rune = '中'
	fmt.Printf("rune: %c, 值: %d\n", r, r)

	// 字符串转换
	s := "Hello,世界"
	bytes := []byte(s) // 转为字节切片
	runes := []rune(s) // 转为rune切片

	fmt.Println("字节切片长度:", len(bytes))   // 12
	fmt.Println("rune切片长度:", len(runes)) // 8
}

func TestMain9(t *testing.T) {
	text := "  Go语言编程入门  "

	fmt.Println("===== 字符串处理 =====")
	fmt.Printf("原始: [%s]\n", text)

	// 去除首尾空格
	trimmed := strings.TrimSpace(text)
	fmt.Printf("去空格: [%s]\n", trimmed)

	// 字符串包含检查
	fmt.Println("包含'语言':", strings.Contains(trimmed, "语言"))

	// 字符串替换
	replaced := strings.Replace(trimmed, "Go", "Golang", 1)
	fmt.Println("替换后:", replaced)

	// 字符串分割
	parts := strings.Split("a,b,c,d", ",")
	fmt.Println("分割结果:", parts)

	// 字符串连接
	joined := strings.Join(parts, "-")
	fmt.Println("连接结果:", joined)

	// 大小写转换
	fmt.Println("大写:", strings.ToUpper("hello"))
	fmt.Println("小写:", strings.ToLower("WORLD"))

	// 统计字符数
	fmt.Println("字符数:", utf8.RuneCountInString(trimmed))
}
