package main14

import (
	"fmt"
	"strconv"
)

type Config struct {
	Port    int
	Debug   bool
	Timeout float64
}

func parseArg(key, value string, config *Config) error {
	switch key {
	case "port":
		port, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("无效的端口号 %q: %w", value, err)
		}
		if port < 1 || port > 65535 {
			return fmt.Errorf("端口号超出范围: %d", port)
		}
		config.Port = port

	case "debug":
		debug, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("无效的debug值 %q: %w", value, err)
		}
		config.Debug = debug

	case "timeout":
		timeout, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("无效的超时值 %q: %w", value, err)
		}
		if timeout < 0 {
			return fmt.Errorf("超时不能为负数: %f", timeout)
		}
		config.Timeout = timeout

	default:
		return fmt.Errorf("未知参数: %s", key)
	}
	return nil
}
