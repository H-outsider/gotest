package main7

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// divmod 同时返回商和余数。
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func getInfo() (string, int, string) {
	return "张三", 25, "zhangsan@example.com"
}

func minMax(numbers []int) (min int, max int) {
	if len(numbers) == 0 {
		return
	}

	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return
}

func findUser(id int) (string, bool) {
	users := map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}

	name, exists := users[id]
	return name, exists
}

// parseAddress 解析 "host:port" 格式的地址。
func parseAddress(addr string) (host string, port int, err error) {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		err = fmt.Errorf("无效的地址格式: %s", addr)
		return
	}

	host = parts[0]
	port, err = strconv.Atoi(parts[1])
	if err != nil {
		err = fmt.Errorf("无效的端口号: %s", parts[1])
		return
	}

	if port < 1 || port > 65535 {
		err = fmt.Errorf("端口号超出范围: %d", port)
		return
	}

	return
}
