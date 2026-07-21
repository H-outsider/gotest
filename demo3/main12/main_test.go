package main12

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	t.Skip("演示用例：直接panic会导致测试失败")
	panicBasic()
}

func TestMain2(t *testing.T) {
	t.Skip("演示用例：panic前会先执行defer，但最终仍会导致测试失败")
	panicWithDefer()
}

func TestMain3(t *testing.T) {
	safeCall()
	fmt.Println("程序继续运行")
}

func TestMain4(t *testing.T) {
	t.Skip("演示用例：recover不在defer中，无法捕获panic")
	recoverOutsideDeferWrong()
}

func TestMain5(t *testing.T) {
	t.Skip("演示用例：recover在嵌套函数中，无法捕获当前panic")
	nestedRecoverWrong()
}

func TestMain6(t *testing.T) {
	recoverPanicType()
}

func TestMain7(t *testing.T) {
	// 包装处理函数
	safeHandler := recoveryMiddleware(userHandler)

	// 正常请求
	safeHandler("/users")

	// 会panic的请求
	safeHandler("/crash")

	// 服务器继续运行
	safeHandler("/products")
}

func TestMain8(t *testing.T) {
	// 正常函数
	ok, msg := safeExecute(func() {
		fmt.Println("正常执行")
	})
	fmt.Printf("结果: success=%v, error=%q\n", ok, msg)

	// 会panic的函数
	ok, msg = safeExecute(func() {
		panic("出错了")
	})
	fmt.Printf("结果: success=%v, error=%q\n", ok, msg)
}
