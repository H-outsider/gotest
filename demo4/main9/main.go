package main9

import (
	"fmt"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age  int
}

// String 实现Stringer接口。
func (p Person) String() string {
	return fmt.Sprintf("%s (%d岁)", p.Name, p.Age)
}

type User struct {
	Username string
	Password string
	Email    string
}

func (u User) String() string {
	// 隐藏密码
	return fmt.Sprintf("User{Username: %s, Email: %s}",
		u.Username, u.Email)
}

type IPAddress [4]byte

func (ip IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type NetworkConfig struct {
	IP      IPAddress
	Subnet  IPAddress
	Gateway IPAddress
}

func (n NetworkConfig) String() string {
	var sb strings.Builder
	sb.WriteString("网络配置:\n")
	sb.WriteString(fmt.Sprintf("  IP地址: %s\n", n.IP))
	sb.WriteString(fmt.Sprintf("  子网掩码: %s\n", n.Subnet))
	sb.WriteString(fmt.Sprintf("  网关: %s", n.Gateway))
	return sb.String()
}

type Status int

const (
	Pending Status = iota
	Running
	Completed
	Failed
)

func (s Status) String() string {
	names := []string{"待处理", "运行中", "已完成", "失败"}
	if int(s) < len(names) {
		return names[s]
	}
	return fmt.Sprintf("未知状态(%d)", s)
}

type Task struct {
	Name   string
	Status Status
}

func (t Task) String() string {
	return fmt.Sprintf("任务[%s]: %s", t.Name, t.Status)
}

type Counter struct {
	Value int
}

// String 使用指针接收者。
func (c *Counter) String() string {
	if c == nil {
		return "Counter(nil)"
	}
	return fmt.Sprintf("Counter(%d)", c.Value)
}

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (l LogLevel) String() string {
	return []string{"DEBUG", "INFO", "WARN", "ERROR"}[l]
}

type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
	Fields    map[string]interface{}
}

func (e LogEntry) String() string {
	var sb strings.Builder

	// 时间戳
	sb.WriteString(e.Timestamp.Format("2006-01-02 15:04:05"))
	sb.WriteString(" ")

	// 日志级别
	sb.WriteString(fmt.Sprintf("[%-5s] ", e.Level))

	// 消息
	sb.WriteString(e.Message)

	// 附加字段
	if len(e.Fields) > 0 {
		sb.WriteString(" {")
		first := true
		for k, v := range e.Fields {
			if !first {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s=%v", k, v))
			first = false
		}
		sb.WriteString("}")
	}

	return sb.String()
}
