package main8

import (
	"fmt"
	"sort"
	"testing"
)

func TestMain1(t *testing.T) {
	// 整数切片排序
	ints := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(ints)
	fmt.Println("整数排序:", ints)

	// 字符串切片排序
	strs := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(strs)
	fmt.Println("字符串排序:", strs)

	// 浮点数切片排序
	floats := []float64{3.14, 1.41, 2.71, 0.58}
	sort.Float64s(floats)
	fmt.Println("浮点数排序:", floats)

	// 逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("逆序排序:", ints)
}

func TestMain2(t *testing.T) {
	people := []Person{
		{"张三", 25},
		{"李四", 30},
		{"王五", 20},
		{"赵六", 28},
	}

	// 按年龄排序
	sort.Sort(ByAge(people))
	fmt.Println("按年龄排序:")
	for _, p := range people {
		fmt.Printf("  %s: %d岁\n", p.Name, p.Age)
	}

	// 按姓名排序
	sort.Sort(ByName(people))
	fmt.Println("\n按姓名排序:")
	for _, p := range people {
		fmt.Printf("  %s: %d岁\n", p.Name, p.Age)
	}
}

func TestMain3(t *testing.T) {
	people := []Person{
		{"张三", 25},
		{"李四", 30},
		{"王五", 20},
		{"赵六", 28},
	}

	// 按年龄排序（更简洁）
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("按年龄排序:", people)

	// 按姓名逆序排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})
	fmt.Println("按姓名逆序:", people)
}

func TestMain4(t *testing.T) {
	items := []Item{
		{"A", 1, 1},
		{"B", 2, 2},
		{"C", 1, 3},
		{"D", 2, 4},
		{"E", 1, 5},
	}

	// 稳定排序：相同Value的元素保持原始顺序
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	})

	fmt.Println("稳定排序后:")
	for _, item := range items {
		fmt.Printf("  %s (Value=%d, 原序=%d)\n",
			item.Name, item.Value, item.Order)
	}
}

func TestMain5(t *testing.T) {
	students := []Student{
		{"张三", "A班", 85},
		{"李四", "B班", 92},
		{"王五", "A班", 78},
		{"赵六", "B班", 85},
		{"钱七", "A班", 92},
		{"孙八", "B班", 78},
	}

	// 先按班级，班级相同按分数降序，分数相同按姓名
	sort.Slice(students, func(i, j int) bool {
		if students[i].Class != students[j].Class {
			return students[i].Class < students[j].Class
		}
		if students[i].Score != students[j].Score {
			return students[i].Score > students[j].Score // 降序
		}
		return students[i].Name < students[j].Name
	})

	fmt.Println("成绩单（按班级、分数、姓名排序）:")
	currentClass := ""
	for _, s := range students {
		if s.Class != currentClass {
			currentClass = s.Class
			fmt.Printf("\n%s:\n", currentClass)
		}
		fmt.Printf("  %s: %d分\n", s.Name, s.Score)
	}
}

func TestMain6(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}

	// 查找第一个 >= target 的位置
	target := 6
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	if i < len(nums) && nums[i] == target {
		fmt.Printf("找到 %d，位置: %d\n", target, i)
	} else {
		fmt.Printf("%d 不存在，应插入位置: %d\n", target, i)
	}

	// 查找存在的值
	target = 7
	i = sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	fmt.Printf("找到 %d，位置: %d\n", target, i)
}
