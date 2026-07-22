package main6

import (
	"fmt"
	"strings"
	"time"
)

func describeValue(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整数: %d, 平方: %d\n", v, v*v)
	case string:
		fmt.Printf("字符串: %q, 长度: %d\n", v, len(v))
	case bool:
		fmt.Printf("布尔值: %v\n", v)
	case float64:
		fmt.Printf("浮点数: %.2f\n", v)
	default:
		fmt.Printf("未知类型: %T, 值: %v\n", v, v)
	}
}

func printNumber(i interface{}) {
	switch v := i.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("有符号整数:", v)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Println("无符号整数:", v)
	case float32, float64:
		fmt.Println("浮点数:", v)
	default:
		fmt.Println("不是数字:", v)
	}
}

type Stringer interface {
	String() string
}

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) String() string {
	return "Dog: " + d.Name
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says woof!")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "says meow!")
}

func process(i interface{}) {
	switch v := i.(type) {
	case Stringer:
		fmt.Println("实现了Stringer:", v.String())
	case Speaker:
		fmt.Print("实现了Speaker: ")
		v.Speak()
	default:
		fmt.Printf("普通类型: %T\n", v)
	}
}

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func handleNil(i interface{}) {
	switch v := i.(type) {
	case nil:
		fmt.Println("值是nil")
	case int:
		fmt.Println("整数:", v)
	default:
		fmt.Printf("类型: %T\n", v)
	}
}

func format(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return "<nil>"
	case bool:
		if v {
			return "是"
		}
		return "否"
	case int, int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	case string:
		if v == "" {
			return "<空字符串>"
		}
		return v
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case []string:
		return "[" + strings.Join(v, ", ") + "]"
	case []int:
		strs := make([]string, len(v))
		for i, n := range v {
			strs[i] = fmt.Sprintf("%d", n)
		}
		return "[" + strings.Join(strs, ", ") + "]"
	case error:
		return "错误: " + v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("用户(%s, %d岁)", u.Name, u.Age)
}
