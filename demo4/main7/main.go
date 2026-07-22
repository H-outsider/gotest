package main7

import (
	"fmt"
	"strconv"
)

// printValue 接受任意类型的参数。
func printValue(v any) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

// printAll 接受任意数量任意类型的参数（类似fmt.Println）。
func printAll(values ...any) {
	for i, v := range values {
		fmt.Printf("[%d] %v (%T)\n", i, v, v)
	}
}

// maxAny 使用空接口（运行时类型检查）。
func maxAny(a, b any) any {
	switch va := a.(type) {
	case int:
		if vb, ok := b.(int); ok && va > vb {
			return va
		}
		return b
	case float64:
		if vb, ok := b.(float64); ok && va > vb {
			return va
		}
		return b
	default:
		return a
	}
}

// maxGeneric 使用泛型（编译时类型检查，Go 1.18+）。
func maxGeneric[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Config struct {
	data map[string]any
}

func NewConfig() *Config {
	return &Config{
		data: make(map[string]any),
	}
}

func (c *Config) Set(key string, value any) {
	c.data[key] = value
}

func (c *Config) Get(key string) (any, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *Config) GetString(key string) string {
	if v, ok := c.data[key]; ok {
		switch val := v.(type) {
		case string:
			return val
		case int:
			return strconv.Itoa(val)
		case bool:
			return strconv.FormatBool(val)
		default:
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

func (c *Config) GetInt(key string) int {
	if v, ok := c.data[key]; ok {
		switch val := v.(type) {
		case int:
			return val
		case float64:
			return int(val)
		case string:
			n, _ := strconv.Atoi(val)
			return n
		}
	}
	return 0
}

func (c *Config) GetBool(key string) bool {
	if v, ok := c.data[key]; ok {
		switch val := v.(type) {
		case bool:
			return val
		case int:
			return val != 0
		case string:
			return val == "true" || val == "1"
		}
	}
	return false
}
