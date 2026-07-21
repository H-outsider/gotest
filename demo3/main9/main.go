package main9

import "fmt"

func counter() func() int {
	count := 0 // 这个变量会被闭包"捕获"
	return func() int {
		count++
		return count
	}
}

// mapInts 对切片中的每个元素应用函数。
func mapInts(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = f(n)
	}
	return result
}

// makeAdder 创建一个加法器。
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// makeMultiplier 创建一个乘法器。
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

// memoize 创建带缓存的计算函数。
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if result, ok := cache[n]; ok {
			fmt.Printf("缓存命中: %d\n", n)
			return result
		}
		result := f(n)
		cache[n] = result
		fmt.Printf("计算并缓存: %d -> %d\n", n, result)
		return result
	}
}

func expensiveCalculation(n int) int {
	// 模拟耗时计算
	result := n * n
	return result
}
