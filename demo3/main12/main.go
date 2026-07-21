package main12

import (
	"fmt"
	"log"
)

func panicBasic() {
	fmt.Println("开始")
	panic("出大事了！")
	fmt.Println("这行不会执行")
}

func panicWithDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("正常执行")
	panic("发生panic")
	fmt.Println("panic后不会执行")
}

func mightPanic() {
	panic("出错了")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到panic:", r)
		}
	}()

	mightPanic()
	fmt.Println("这行不会执行")
}

func recoverOutsideDeferWrong() {
	// 错误用法：recover不在defer中
	if r := recover(); r != nil {
		fmt.Println("捕获:", r) // 不会执行
	}

	panic("test")
}

func nestedRecoverWrong() {
	// 错误用法：recover在嵌套函数中
	defer func() {
		nestedRecover() // 无法捕获
	}()

	panic("test")
}

func nestedRecover() {
	if r := recover(); r != nil {
		fmt.Println("捕获:", r) // 不会执行
	}
}

func recoverPanicType() {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case string:
				fmt.Println("字符串panic:", v)
			case error:
				fmt.Println("error类型panic:", v.Error())
			default:
				fmt.Printf("未知类型panic: %T %v\n", v, v)
			}
		}
	}()

	// 可以panic任何值
	// panic("字符串")
	// panic(errors.New("error对象"))
	panic(42)
}

// HandlerFunc 模拟HTTP处理函数。
type HandlerFunc func(path string)

// recoveryMiddleware 恢复中间件。
func recoveryMiddleware(next HandlerFunc) HandlerFunc {
	return func(path string) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("请求 %s 发生panic: %v", path, r)
				fmt.Println("返回500错误给客户端")
			}
		}()
		next(path)
	}
}

// userHandler 可能panic的处理函数。
func userHandler(path string) {
	if path == "/crash" {
		panic("模拟崩溃")
	}
	fmt.Println("处理请求:", path)
}

// safeExecute 安全执行函数，返回是否成功和错误信息。
func safeExecute(f func()) (success bool, errMsg string) {
	defer func() {
		if r := recover(); r != nil {
			success = false
			errMsg = fmt.Sprintf("%v", r)
		}
	}()

	f()
	return true, ""
}
