package main10

import (
	"errors"
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Println("结果:", result)
}

func TestMain2(t *testing.T) {
	err := openFile("config.txt")
	if err != nil {
		fmt.Println("错误:", err)
	}
}

func TestMain3(t *testing.T) {
	err := validateAge(-5)
	if err != nil {
		fmt.Println(err)

		// 类型断言获取详细信息
		if ve, ok := err.(*ValidationError); ok {
			fmt.Println("错误字段:", ve.Field)
			fmt.Println("错误信息:", ve.Message)
		}
	}
}

func TestMain4(t *testing.T) {
	err := fetchData("https://api.example.com/users")
	if err != nil {
		fmt.Println("请求失败:", err)

		if httpErr, ok := err.(*HTTPError); ok {
			if httpErr.IsNotFound() {
				fmt.Println("资源不存在")
			}
			if httpErr.IsServerError() {
				fmt.Println("服务器错误，请稍后重试")
			}
		}
	}
}

func TestMain5(t *testing.T) {
	_, err := getUser(0)

	// 使用==比较哨兵错误
	if err == ErrNotFound {
		fmt.Println("用户不存在")
	} else if err == ErrUnauthorized {
		fmt.Println("没有权限")
	}

	// 或使用errors.Is（推荐，支持错误链）
	if errors.Is(err, ErrNotFound) {
		fmt.Println("使用errors.Is: 用户不存在")
	}
}

func TestMain6(t *testing.T) {
	testCases := []int{-1, 50, 200}

	for _, id := range testCases {
		user, err := findUserWithAppError(id)
		if err != nil {
			fmt.Printf("查询ID=%d失败: %v\n", id, err)

			// 检查错误类型
			var appErr *AppError
			if errors.As(err, &appErr) {
				fmt.Printf("  错误码: %s\n", appErr.Code)
			}

			// 检查是否包含特定错误
			if errors.Is(err, ErrUserNotFound) {
				fmt.Println("  (用户确实不存在)")
			}
			continue
		}
		fmt.Printf("查询ID=%d成功: %s\n", id, user)
	}
}
