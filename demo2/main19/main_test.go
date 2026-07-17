package main19

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	// 声明nil切片
	var s1 []int
	fmt.Println("nil切片:", s1, "是否为nil:", s1 == nil)

	// 字面量初始化
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println("字面量:", s2)

	// 使用make创建
	s3 := make([]int, 3) // 长度为3，元素为零值
	fmt.Println("make创建:", s3)

	// 使用make创建，指定容量
	s4 := make([]int, 3, 10) // 长度3，容量10
	fmt.Println("指定容量:", s4, "len:", len(s4), "cap:", cap(s4))
}

func TestMain2(t *testing.T) {
	s := []string{"苹果", "香蕉", "橙子"}

	// 读取
	fmt.Println("第一个:", s[0])
	fmt.Println("最后一个:", s[len(s)-1])

	// 修改
	s[1] = "葡萄"
	fmt.Println("修改后:", s)

	// 长度和容量
	fmt.Println("长度:", len(s))
	fmt.Println("容量:", cap(s))
}

func TestMain3(t *testing.T) {
	arr := [5]int{10, 20, 30, 40, 50}

	// 从数组创建切片
	s1 := arr[1:4] // 包含索引1,2,3，不包含4
	fmt.Println("arr[1:4]:", s1)

	s2 := arr[:3] // 从开头到索引2
	fmt.Println("arr[:3]:", s2)

	s3 := arr[2:] // 从索引2到末尾
	fmt.Println("arr[2:]:", s3)

	s4 := arr[:] // 整个数组
	fmt.Println("arr[:]:", s4)
}

func TestMain4(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}

	s1 := arr[1:4]
	s2 := arr[2:5]

	fmt.Println("原始 arr:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// 修改s1会影响arr和s2
	s1[1] = 300

	fmt.Println("\n修改s1[1]=300后:")
	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}

func TestMain5(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}

	// for-range（推荐）
	fmt.Println("for-range遍历:")
	for i, v := range nums {
		fmt.Printf("索引%d: %d\n", i, v)
	}

	// 普通for循环
	fmt.Println("\n普通for循环:")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("索引%d: %d\n", i, nums[i])
	}
}

func TestMain6(t *testing.T) {
	s := make([]int, 3, 10)

	fmt.Println("切片:", s)
	fmt.Println("长度 len(s):", len(s)) // 元素个数
	fmt.Println("容量 cap(s):", cap(s)) // 底层数组的可用空间

	// 从原切片创建子切片
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sub := arr[2:5]

	fmt.Println("\narr:", arr)
	fmt.Println("sub = arr[2:5]:", sub)
	fmt.Println("len(sub):", len(sub)) // 3
	fmt.Println("cap(sub):", cap(sub)) // 8（从索引2到末尾）
}

func TestMain7(t *testing.T) {
	// 模拟动态收集用户ID
	var userIDs []int

	// 模拟添加用户
	for i := 1; i <= 5; i++ {
		userIDs = append(userIDs, i*100)
		fmt.Printf("添加用户%d后: %v, len=%d, cap=%d\n",
			i, userIDs, len(userIDs), cap(userIDs))
	}

	fmt.Println("\n最终用户列表:", userIDs)
}
