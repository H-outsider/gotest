package main18

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	// 声明一个包含5个int的数组，元素自动初始化为零值
	var nums [5]int
	fmt.Println("零值数组:", nums)

	// 声明并初始化
	var scores [3]int = [3]int{85, 92, 78}
	fmt.Println("初始化数组:", scores)

	// 简短声明
	names := [3]string{"张三", "李四", "王五"}
	fmt.Println("字符串数组:", names)
}

func TestMain2(t *testing.T) {
	// 编译器会推断长度为5
	nums := [...]int{10, 20, 30, 40, 50}
	fmt.Println("数组:", nums)
	fmt.Println("长度:", len(nums))
}

func TestMain3(t *testing.T) {
	// 指定索引位置初始化
	nums := [5]int{0: 100, 3: 400}
	fmt.Println(nums) // [100 0 0 400 0]

	// 混合使用
	mixed := [...]int{1, 2, 4: 5, 6}
	fmt.Println(mixed) // [1 2 0 0 5 6]，长度为6
}

func TestMain4(t *testing.T) {
	nums := [5]int{10, 20, 30, 40, 50}

	// 读取元素（索引从0开始）
	fmt.Println("第一个元素:", nums[0])
	fmt.Println("最后一个元素:", nums[4])

	// 修改元素
	nums[2] = 300
	fmt.Println("修改后:", nums)

	// 获取长度
	fmt.Println("数组长度:", len(nums))
}

func TestMain5(t *testing.T) {
	fruits := [4]string{"苹果", "香蕉", "橙子", "葡萄"}

	// 方式1：普通for循环
	fmt.Println("普通for循环:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  索引%d: %s\n", i, fruits[i])
	}

	// 方式2：for-range（推荐）
	fmt.Println("\nfor-range:")
	for i, fruit := range fruits {
		fmt.Printf("  索引%d: %s\n", i, fruit)
	}

	// 方式3：只要值
	fmt.Println("\n只要值:")
	for _, fruit := range fruits {
		fmt.Println(" ", fruit)
	}
}

func TestMain6(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a // 复制整个数组

	b[0] = 100 // 修改b

	fmt.Println("a:", a) // a不变
	fmt.Println("b:", b) // b变了
}

func TestMain7(t *testing.T) {
	// 二维数组：3行4列
	var matrix [3][4]int

	// 初始化
	matrix[0][0] = 1
	matrix[1][1] = 5
	matrix[2][3] = 9

	fmt.Println("二维数组:")
	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}

	// 声明并初始化二维数组
	grid := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n初始化的二维数组:")
	for _, row := range grid {
		fmt.Println(row)
	}
}

func TestMain8(t *testing.T) {
	scores := [5]int{85, 92, 78, 96, 88}

	// 计算总分
	sum := 0
	for _, score := range scores {
		sum += score
	}

	// 计算平均分
	avg := float64(sum) / float64(len(scores))

	// 找最高分和最低分
	max, min := scores[0], scores[0]
	for _, score := range scores {
		if score > max {
			max = score
		}
		if score < min {
			min = score
		}
	}

	fmt.Println("成绩:", scores)
	fmt.Println("总分:", sum)
	fmt.Printf("平均分: %.1f\n", avg)
	fmt.Println("最高分:", max)
	fmt.Println("最低分:", min)
}

func TestMain9(t *testing.T) {
	var a [3]int
	var b [4]int

	fmt.Printf("a的类型: %T\n", a)
	fmt.Printf("b的类型: %T\n", b)

	// a = b  // 编译错误：cannot use b (type [4]int) as type [3]int
}
