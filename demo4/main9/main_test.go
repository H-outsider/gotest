package main9

import (
	"fmt"
	"testing"
	"time"
)

func TestMain1(t *testing.T) {
	p := Person{Name: "张三", Age: 25}

	// 这些都会调用String()方法
	fmt.Println(p)
	fmt.Printf("%v\n", p)
	fmt.Printf("%s\n", p)

	// %+v 和 %#v 不受影响
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
}

func TestMain2(t *testing.T) {
	user := User{
		Username: "admin",
		Password: "secret123",
		Email:    "admin@example.com",
	}

	fmt.Println(user)
	fmt.Printf("用户信息: %v\n", user)
}

func TestMain3(t *testing.T) {
	config := NetworkConfig{
		IP:      IPAddress{192, 168, 1, 100},
		Subnet:  IPAddress{255, 255, 255, 0},
		Gateway: IPAddress{192, 168, 1, 1},
	}

	fmt.Println(config)
}

func TestMain4(t *testing.T) {
	tasks := []Task{
		{"数据备份", Completed},
		{"日志清理", Running},
		{"报表生成", Pending},
		{"系统检查", Failed},
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
}

func TestMain5(t *testing.T) {
	c1 := &Counter{Value: 42}
	fmt.Println(c1)

	var c2 *Counter
	fmt.Println(c2) // 会调用String()，输出Counter(nil)

	c3 := Counter{Value: 10}
	fmt.Println(c3)  // 值类型，不会调用*Counter的String()
	fmt.Println(&c3) // 指针类型，会调用
}

func TestMain6(t *testing.T) {
	entries := []LogEntry{
		{
			Timestamp: time.Now(),
			Level:     INFO,
			Message:   "服务启动",
			Fields:    map[string]interface{}{"port": 8080},
		},
		{
			Timestamp: time.Now(),
			Level:     WARN,
			Message:   "内存使用率过高",
			Fields:    map[string]interface{}{"usage": "85%", "threshold": "80%"},
		},
		{
			Timestamp: time.Now(),
			Level:     ERROR,
			Message:   "数据库连接失败",
			Fields:    map[string]interface{}{"host": "localhost", "retry": 3},
		},
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
}
