package main3

import (
	"fmt"
	"testing"
)

const Pi = 3.14159
const AppVersion = "1.0.0"

const (
	StatusOK       = 200
	StatusNotFound = 404
	StatusError    = 500
)

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const (
	_  = iota             // 0，用_忽略
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1048576
	GB                    // 1 << 30 = 1073741824
	TB                    // 1 << 40
)

const big = 1e100        // 非常大的数
const small = big / 1e99 // 结果是10

func TestMain1(t *testing.T) {
	fmt.Println("圆周率:", Pi)
	fmt.Println("应用版本:", AppVersion)
}

func TestMain2(t *testing.T) {
	fmt.Println("成功:", StatusOK)
	fmt.Println("未找到:", StatusNotFound)
	fmt.Println("错误:", StatusError)
}

func TestMain3(t *testing.T) {
	fmt.Println("星期日:", Sunday)
	fmt.Println("星期一:", Monday)
	fmt.Println("星期六:", Saturday)
}

func TestMain4(t *testing.T) {
	fmt.Println("1 KB =", KB, "字节")
	fmt.Println("1 MB =", MB, "字节")
	fmt.Println("1 GB =", GB, "字节")
}

func TestMain5(t *testing.T) {
	var f float64 = small
	var i int = small

	fmt.Println("float64:", f)
	fmt.Println("int:", i)
}

func TestMain6(t *testing.T) {
	currentState := StateConnected

	fmt.Println("===== 服务器配置 =====")
	fmt.Println("最大连接数:", MaxConnections)
	fmt.Println("超时时间:", Timeout, "秒")
	fmt.Println("重试次数:", RetryCount)
	fmt.Println()
	fmt.Println("当前状态:", getStateName(currentState))
}
