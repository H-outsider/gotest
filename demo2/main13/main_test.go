package main13

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	// i := 0 是初始化
	// i < 5 是条件
	// i++ 是后置语句（每次循环后执行）
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func TestMain2(t *testing.T) {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func TestMain3(t *testing.T) {
	count := 0
	for {
		fmt.Println("循环中...", count)
		count++
		if count >= 3 {
			break // 用break退出循环
		}
	}
	fmt.Println("循环结束")
}

func TestMain4(t *testing.T) {
	// 省略初始化（变量在外部声明）
	i := 0
	for ; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("---")

	// 省略后置语句（在循环体内处理）
	for j := 0; j < 3; {
		fmt.Println(j)
		j += 2
	}
}

func TestMain5(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的和:", sum)

	// 使用公式验证：n*(n+1)/2
	fmt.Println("公式验证:", 100*101/2)
}

func TestMain6(t *testing.T) {
	// 嵌套循环
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func TestMain7(t *testing.T) {
	// 从大到小
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("发射！")
}

func TestMain8(t *testing.T) {
	// 步长为2（只打印偶数）
	fmt.Println("偶数:")
	for i := 0; i <= 10; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 步长为3
	fmt.Println("3的倍数:")
	for i := 0; i <= 15; i += 3 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
