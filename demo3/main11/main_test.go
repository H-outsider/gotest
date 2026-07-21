package main11

import (
	"errors"
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	}
}

func TestMain2(t *testing.T) {
	err := openFile("config.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
}

func TestMain3(t *testing.T) {
	err := validateUser("", 25)
	if err != nil {
		fmt.Println(err)

		// 类型断言获取详细信息
		if ve, ok := err.(*ValidationError); ok {
			fmt.Println("错误字段:", ve.Field)
		}
	}

	err = validateUser("张三", 200)
	if err != nil {
		fmt.Println(err)
	}
}

func TestMain4(t *testing.T) {
	err := findUser(123)
	if err != nil {
		fmt.Println("错误:", err)

		// 检查是否包含特定错误
		if errors.Is(err, ErrNotFound) {
			fmt.Println("这是一个未找到错误")
		}
	}
}

func TestMain5(t *testing.T) {
	err := executeQuery("SELECT * FROM users")
	if err != nil {
		fmt.Println("错误:", err)

		// 检查错误链中是否包含特定错误
		if errors.Is(err, ErrConnection) {
			fmt.Println("发现连接错误，尝试重连...")
		}

		// 获取特定类型的错误
		var qe *QueryError
		if errors.As(err, &qe) {
			fmt.Println("失败的查询:", qe.Query)
		}
	}
}

func TestMain6(t *testing.T) {
	// 正常配置
	content := `
host = localhost
port = 8080
`
	config, err := parseConfig(content)
	if err != nil {
		fmt.Println("加载失败:", err)
	} else {
		fmt.Printf("配置: Host=%s, Port=%d\n", config.Host, config.Port)
	}

	// 无效配置
	badContent := "port = abc"
	_, err = parseConfig(badContent)
	if err != nil {
		fmt.Println("加载失败:", err)
	}

	// 空配置
	_, err = parseConfig("")
	if errors.Is(err, ErrEmptyConfig) {
		fmt.Println("配置文件为空")
	}
}
