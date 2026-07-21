package main10

import (
	"fmt"
	"sync"
	"testing"
)

func TestMain1(t *testing.T) {
	fmt.Println("开始")
	defer fmt.Println("延迟执行")
	fmt.Println("结束")
}

func TestMain2(t *testing.T) {
	defer fmt.Println("第一个defer")
	defer fmt.Println("第二个defer")
	defer fmt.Println("第三个defer")
	fmt.Println("正常代码")
}

func TestMain3(t *testing.T) {
	err := readFile("test.txt")
	if err != nil {
		fmt.Println("错误:", err)
	}
}

func TestMain4(t *testing.T) {
	counter = 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func TestMain5(t *testing.T) {
	x := 10
	defer fmt.Println("defer中的x:", x) // x在defer时就确定了
	x = 20
	fmt.Println("当前x:", x)
}

func TestMain6(t *testing.T) {
	fmt.Println("example1:", example1())
	fmt.Println("example2:", example2())
}

func TestMain7(t *testing.T) {
	slowOperation()
}

func TestMain8(t *testing.T) {
	db := &DB{}

	fmt.Println("=== 成功场景 ===")
	_ = doTransaction(db, false)

	fmt.Println("\n=== 失败场景 ===")
	_ = doTransaction(db, true)
}
