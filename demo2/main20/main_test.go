package main20

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println("原始切片:", s)

	s = append(s, 4)
	fmt.Println("追加4后:", s)

	s = append(s, 5)
	fmt.Println("追加5后:", s)
}

func TestMain2(t *testing.T) {
	s := []int{1, 2, 3}

	// 一次追加多个元素
	s = append(s, 4, 5, 6)
	fmt.Println(s)
}

func TestMain3(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// 将s2的所有元素追加到s1
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func TestMain4(t *testing.T) {
	var s []int // nil切片

	fmt.Println("nil切片:", s, "len:", len(s))

	// 可以直接append，不需要先make
	s = append(s, 1, 2, 3)
	fmt.Println("追加后:", s)
}

func TestMain5(t *testing.T) {
	var s []int

	for i := 1; i <= 10; i++ {
		s = append(s, i)
		fmt.Printf("len=%2d, cap=%2d, %v\n", len(s), cap(s), s)
	}
}

func TestMain6(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	// 复制src到dst
	copied := copy(dst, src)

	fmt.Println("源切片:", src)
	fmt.Println("目标切片:", dst)
	fmt.Println("复制了", copied, "个元素")

	// 修改dst不会影响src
	dst[0] = 100
	fmt.Println("\n修改dst后:")
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
}

func TestMain7(t *testing.T) {
	var todos []string

	// 添加待办事项
	todos = append(todos, "学习Go语言")
	todos = append(todos, "写代码练习")
	todos = append(todos, "阅读文档")

	fmt.Println("===== 待办事项 =====")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}

	// 批量添加
	moreTodos := []string{"提交代码", "代码审查"}
	todos = append(todos, moreTodos...)

	fmt.Println("\n===== 添加后 =====")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}

	// 删除第2项（索引1）
	todos = append(todos[:1], todos[2:]...)

	fmt.Println("\n===== 删除第2项后 =====")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}
}

func TestMain8(t *testing.T) {
	s1 := make([]int, 3, 3) // len=3, cap=3
	s1[0], s1[1], s1[2] = 1, 2, 3

	s2 := append(s1, 4) // 容量不足，创建新底层数组

	s2[0] = 100

	fmt.Println("s1:", s1) // [1 2 3]，未受影响
	fmt.Println("s2:", s2) // [100 2 3 4]
}
