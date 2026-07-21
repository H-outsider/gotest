package main13

import (
	"fmt"
	"strings"
	"testing"
)

func TestMain1(t *testing.T) {
	s := "Hello, World!"

	// 包含
	fmt.Println("包含World:", strings.Contains(s, "World"))
	fmt.Println("包含Go:", strings.Contains(s, "Go"))

	// 前缀和后缀
	fmt.Println("以Hello开头:", strings.HasPrefix(s, "Hello"))
	fmt.Println("以!结尾:", strings.HasSuffix(s, "!"))

	// 查找位置
	fmt.Println("o的位置:", strings.Index(s, "o"))
	fmt.Println("最后一个o的位置:", strings.LastIndex(s, "o"))
	fmt.Println("x的位置:", strings.Index(s, "x")) // -1表示未找到

	// 统计出现次数
	fmt.Println("o出现次数:", strings.Count(s, "o"))
}

func TestMain2(t *testing.T) {
	// 分割
	csv := "apple,banana,orange"
	fruits := strings.Split(csv, ",")
	fmt.Println("分割结果:", fruits)

	// 按空白分割（连续空白视为一个）
	sentence := "hello   world  go"
	words := strings.Fields(sentence)
	fmt.Println("Fields结果:", words)

	// 限制分割次数
	data := "a:b:c:d:e"
	parts := strings.SplitN(data, ":", 3)
	fmt.Println("SplitN结果:", parts) // [a b c:d:e]

	// 拼接
	items := []string{"Go", "Python", "Java"}
	joined := strings.Join(items, " | ")
	fmt.Println("拼接结果:", joined)
}

func TestMain3(t *testing.T) {
	s := "hello hello hello"

	// 替换指定次数
	fmt.Println("替换2次:", strings.Replace(s, "hello", "hi", 2))

	// 替换所有（-1表示全部替换）
	fmt.Println("替换全部:", strings.Replace(s, "hello", "hi", -1))

	// ReplaceAll（推荐用于全部替换）
	fmt.Println("ReplaceAll:", strings.ReplaceAll(s, "hello", "hi"))

	// 多个替换
	replacer := strings.NewReplacer(
		"hello", "hi",
		" ", "_",
	)
	fmt.Println("多个替换:", replacer.Replace(s))
}

func TestMain4(t *testing.T) {
	s := "Hello World"

	fmt.Println("大写:", strings.ToUpper(s))
	fmt.Println("小写:", strings.ToLower(s))
	fmt.Println("标题格式:", strings.Title(s)) // 每个单词首字母大写（已弃用）

	// 忽略大小写比较
	fmt.Println("忽略大小写比较:", strings.EqualFold("Go", "go"))
}

func TestMain5(t *testing.T) {
	s := "  hello world  "

	// 去除首尾空白
	fmt.Printf("TrimSpace: %q\n", strings.TrimSpace(s))

	// 去除指定字符
	s2 := "###hello###"
	fmt.Printf("Trim: %q\n", strings.Trim(s2, "#"))

	// 只去除左边或右边
	fmt.Printf("TrimLeft: %q\n", strings.TrimLeft(s2, "#"))
	fmt.Printf("TrimRight: %q\n", strings.TrimRight(s2, "#"))

	// 去除前缀/后缀
	s3 := "hello.txt"
	fmt.Printf("TrimSuffix: %q\n", strings.TrimSuffix(s3, ".txt"))
	fmt.Printf("TrimPrefix: %q\n", strings.TrimPrefix(s3, "hello"))
}

func TestMain6(t *testing.T) {
	configContent := `
# 数据库配置
db_host = localhost
db_port = 3306
db_name = myapp

# 应用配置
app_name = MyApplication
debug = true
`

	config := parseConfig(configContent)

	fmt.Println("解析结果:")
	for key, value := range config {
		fmt.Printf("  %s = %s\n", key, value)
	}
}

func TestMain7(t *testing.T) {
	// 低效方式（每次创建新字符串）
	s := ""
	for i := 0; i < 5; i++ {
		s += fmt.Sprintf("item%d ", i)
	}
	fmt.Println("直接拼接:", s)

	// 高效方式（使用Builder）
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString(fmt.Sprintf("item%d ", i))
	}
	fmt.Println("Builder拼接:", builder.String())
}
