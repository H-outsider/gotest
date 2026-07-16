package main3

// 服务器状态
type ServerState int

const (
	StateIdle      ServerState = iota // 空闲
	StateConnected                    // 已连接
	StateError                        // 错误
	StateRetrying                     // 重试中
)

// 配置常量
const (
	MaxConnections = 100
	Timeout        = 30 // 秒
	RetryCount     = 3
)

func getStateName(state ServerState) string {
	switch state {
	case StateIdle:
		return "空闲"
	case StateConnected:
		return "已连接"
	case StateError:
		return "错误"
	case StateRetrying:
		return "重试中"
	default:
		return "未知"
	}
}
