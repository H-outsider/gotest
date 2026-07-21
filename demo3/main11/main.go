package main11

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrNotFound      = errors.New("未找到")
	ErrConnection    = errors.New("数据库连接失败")
	ErrEmptyConfig   = errors.New("配置为空")
	ErrInvalidFormat = errors.New("配置格式无效")
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func openFile(name string) error {
	if name == "" {
		return fmt.Errorf("文件名不能为空")
	}
	if len(name) > 100 {
		return fmt.Errorf("文件名过长: %d 字符", len(name))
	}
	// 模拟文件不存在
	return fmt.Errorf("文件不存在: %s", name)
}

// ValidationError 自定义错误类型。
type ValidationError struct {
	Field   string
	Message string
}

// Error 实现error接口。
func (e *ValidationError) Error() string {
	return fmt.Sprintf("验证失败 [%s]: %s", e.Field, e.Message)
}

func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "姓名不能为空"}
	}
	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Message: "年龄必须在0-150之间"}
	}
	return nil
}

func findUser(id int) error {
	// 模拟未找到
	return fmt.Errorf("查找用户 %d: %w", id, ErrNotFound)
}

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("查询失败 [%s]: %v", e.Query, e.Err)
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

func executeQuery(query string) error {
	// 模拟连接错误
	return &QueryError{
		Query: query,
		Err:   ErrConnection,
	}
}

type Config struct {
	Host string
	Port int
}

func parseConfig(content string) (*Config, error) {
	if content == "" {
		return nil, ErrEmptyConfig
	}

	lines := strings.Split(content, "\n")
	config := &Config{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("解析行 %q: %w", line, ErrInvalidFormat)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "host":
			config.Host = value
		case "port":
			port, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("解析端口 %q: %w", value, err)
			}
			config.Port = port
		}
	}

	return config, nil
}
