package main13

import "strings"

func parseConfig(content string) map[string]string {
	config := make(map[string]string)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		// 去除空白
		line = strings.TrimSpace(line)

		// 跳过空行和注释
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 分割键值
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		}
	}

	return config
}
