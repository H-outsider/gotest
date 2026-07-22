package main10

import (
	"errors"
	"fmt"
	"time"
)

// 定义哨兵错误（包级变量）
var (
	ErrNotFound     = errors.New("未找到")
	ErrUnauthorized = errors.New("未授权")
	ErrForbidden    = errors.New("禁止访问")
	ErrUserNotFound = errors.New("用户不存在")
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
	// 带格式化的错误信息
	return fmt.Errorf("文件不存在: %s", name)
}

// ValidationError 自定义错误类型。
type ValidationError struct {
	Field   string
	Message string
}

// Error 实现error接口。
func (e *ValidationError) Error() string {
	return fmt.Sprintf("验证错误 [%s]: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "年龄不能为负数",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "年龄不能超过150",
		}
	}
	return nil
}

type HTTPError struct {
	StatusCode int
	URL        string
	Method     string
	Timestamp  time.Time
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("[%s] %s %s 返回 %d",
		e.Timestamp.Format("15:04:05"),
		e.Method,
		e.URL,
		e.StatusCode)
}

// IsNotFound 便于判断的辅助方法。
func (e *HTTPError) IsNotFound() bool {
	return e.StatusCode == 404
}

func (e *HTTPError) IsServerError() bool {
	return e.StatusCode >= 500
}

func fetchData(url string) error {
	// 模拟HTTP请求失败
	return &HTTPError{
		StatusCode: 404,
		URL:        url,
		Method:     "GET",
		Timestamp:  time.Now(),
	}
}

func getUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrNotFound
	}
	if id == 999 {
		return "", ErrUnauthorized
	}
	return "用户" + fmt.Sprint(id), nil
}

// ErrorCode 错误码。
type ErrorCode int

const (
	ErrCodeUnknown ErrorCode = iota
	ErrCodeValidation
	ErrCodeNotFound
	ErrCodeDatabase
	ErrCodeNetwork
)

func (c ErrorCode) String() string {
	names := []string{"UNKNOWN", "VALIDATION", "NOT_FOUND", "DATABASE", "NETWORK"}
	if int(c) < len(names) {
		return names[c]
	}
	return "UNKNOWN"
}

// AppError 应用错误。
type AppError struct {
	Code    ErrorCode
	Message string
	Cause   error // 原始错误
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap 实现Unwrap以支持errors.Is和errors.As。
func (e *AppError) Unwrap() error {
	return e.Cause
}

// NewValidationError 错误构造函数。
func NewValidationError(message string) *AppError {
	return &AppError{Code: ErrCodeValidation, Message: message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: ErrCodeNotFound, Message: message}
}

func WrapError(code ErrorCode, message string, cause error) *AppError {
	return &AppError{Code: code, Message: message, Cause: cause}
}

// findUserWithAppError 业务函数。
func findUserWithAppError(id int) (string, error) {
	if id <= 0 {
		return "", NewValidationError("用户ID必须为正数")
	}
	if id > 100 {
		return "", WrapError(ErrCodeNotFound, "查询用户失败", ErrUserNotFound)
	}
	return fmt.Sprintf("用户%d", id), nil
}
