package main8

import (
	"fmt"
	"time"
)

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func printInfo(values ...int) {
	fmt.Printf("类型: %T\n", values)
	fmt.Printf("长度: %d\n", len(values))
	fmt.Printf("内容: %v\n", values)
}

// join 第一个参数是分隔符，后面是要连接的字符串。
func join(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += sep + parts[i]
	}
	return result
}

func printAll(values ...interface{}) {
	for i, v := range values {
		fmt.Printf("[%d] %T: %v\n", i, v, v)
	}
}

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

func log(level LogLevel, format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] [%s] %s\n", timestamp, level, message)
}
