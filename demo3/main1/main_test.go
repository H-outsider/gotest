package main1

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 截取索引2到5（不含5）
	s1 := s[2:5]
	fmt.Println("s[2:5]:", s1)

	// 从开头到索引4
	s2 := s[:5]
	fmt.Println("s[:5]:", s2)

	// 从索引5到末尾
	s3 := s[5:]
	fmt.Println("s[5:]:", s3)

	// 完整切片
	s4 := s[:]
	fmt.Println("s[:]:", s4)
}

func TestMain2(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原切片: len=%d, cap=%d\n", len(s), cap(s))

	// 从索引2开始截取
	s1 := s[2:5]
	fmt.Printf("s[2:5]: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	// 从索引7开始截取
	s2 := s[7:9]
	fmt.Printf("s[7:9]: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
}

func TestMain3(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	sub := original[1:4]

	fmt.Println("修改前:")
	fmt.Println("original:", original)
	fmt.Println("sub:", sub)

	// 修改sub会影响original
	sub[0] = 100

	fmt.Println("\n修改sub[0]=100后:")
	fmt.Println("original:", original)
	fmt.Println("sub:", sub)

	// 修改original也会影响sub
	original[2] = 200

	fmt.Println("\n修改original[2]=200后:")
	fmt.Println("original:", original)
	fmt.Println("sub:", sub)
}

func TestMain4(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 普通截取：容量到底层数组末尾
	s1 := s[2:5]
	fmt.Printf("s[2:5]: len=%d, cap=%d\n", len(s1), cap(s1))

	// 完整切片表达式：限制容量为6-2=4
	s2 := s[2:5:6]
	fmt.Printf("s[2:5:6]: len=%d, cap=%d\n", len(s2), cap(s2))

	// 完整切片表达式：容量等于长度
	s3 := s[2:5:5]
	fmt.Printf("s[2:5:5]: len=%d, cap=%d\n", len(s3), cap(s3))
}

func TestMain5(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}

	// 方法1：使用copy
	independent := make([]int, 3)
	copy(independent, original[1:4])

	// 方法2：append到nil切片
	independent2 := append([]int(nil), original[1:4]...)

	// 修改不会影响原切片
	independent[0] = 100
	independent2[0] = 200

	fmt.Println("original:", original)
	fmt.Println("independent:", independent)
	fmt.Println("independent2:", independent2)
}

// paginate 返回独立的切片，不共享底层数组。
func paginate(data []int, page, pageSize int) []int {
	start := page * pageSize
	if start >= len(data) {
		return nil
	}

	end := start + pageSize
	if end > len(data) {
		end = len(data)
	}

	// 创建独立副本，避免共享
	result := make([]int, end-start)
	copy(result, data[start:end])
	return result
}

func TestMain6(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	page0 := paginate(data, 0, 3)
	page1 := paginate(data, 1, 3)
	page2 := paginate(data, 2, 3)
	page3 := paginate(data, 3, 3)

	fmt.Println("第0页:", page0)
	fmt.Println("第1页:", page1)
	fmt.Println("第2页:", page2)
	fmt.Println("第3页:", page3)

	// 修改分页结果不影响原数据
	page0[0] = 100
	fmt.Println("\n修改page0后，原数据:", data)
}
