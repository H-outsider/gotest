package main9

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	// 创建一个计数器
	c1 := counter()
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	// 创建另一个计数器，独立的状态
	c2 := counter()
	fmt.Println(c2()) // 1
	fmt.Println(c1()) // 4，c1继续计数
}

func TestMain2(t *testing.T) {
	// 定义并立即调用
	result := func(a, b int) int {
		return a + b
	}(3, 5)
	fmt.Println("立即调用结果:", result)

	// 赋值给变量后调用
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("4 * 5 =", multiply(4, 5))
}

func TestMain3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	// 使用匿名函数作为回调
	doubled := mapInts(nums, func(n int) int {
		return n * 2
	})
	fmt.Println("翻倍:", doubled)

	// 另一个回调
	squared := mapInts(nums, func(n int) int {
		return n * n
	})
	fmt.Println("平方:", squared)
}

func TestMain4(t *testing.T) {
	add5 := makeAdder(5)
	add10 := makeAdder(10)

	fmt.Println("3 + 5 =", add5(3))
	fmt.Println("3 + 10 =", add10(3))

	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("7 * 2 =", double(7))
	fmt.Println("7 * 3 =", triple(7))
}

func TestMain5(t *testing.T) {
	fib := fibonacci()

	fmt.Println("斐波那契数列前10个:")
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()
}

func TestMain6(t *testing.T) {
	cachedCalc := memoize(expensiveCalculation)

	fmt.Println("第一次调用:")
	fmt.Println("square(5) =", cachedCalc(5))
	fmt.Println("square(3) =", cachedCalc(3))

	fmt.Println("\n第二次调用（应该使用缓存）:")
	fmt.Println("square(5) =", cachedCalc(5))
	fmt.Println("square(3) =", cachedCalc(3))

	fmt.Println("\n新的值:")
	fmt.Println("square(7) =", cachedCalc(7))
}
