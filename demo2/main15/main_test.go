package main15

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("遇到5，终止循环")
			break
		}
		fmt.Println(i)
	}
	fmt.Println("循环结束")
}

func TestMain2(t *testing.T) {
	fmt.Println("跳过偶数，只打印奇数:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Println(i)
	}
}

func TestMain3(t *testing.T) {
	count := 0
	for {
		count++
		fmt.Println("第", count, "次循环")

		if count >= 3 {
			break // 达到条件后退出无限循环
		}
	}
	fmt.Println("已退出循环")
}

func TestMain4(t *testing.T) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("外层循环 i=%d\n", i)
		for j := 1; j <= 3; j++ {
			if j == 2 {
				break // 只跳出内层循环
			}
			fmt.Printf("  内层循环 j=%d\n", j)
		}
	}
}

func TestMain5(t *testing.T) {
OuterLoop: // 标签
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 2 && j == 2 {
				fmt.Println("跳出外层循环")
				break OuterLoop // 跳出标签指定的循环
			}
		}
	}
	fmt.Println("循环结束")
}

func TestMain6(t *testing.T) {
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue OuterLoop // 跳到外层循环的下一次迭代
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}

func TestMain7(t *testing.T) {
	numbers := []int{23, 45, 67, 89, 12, 34, 56}
	target := 89

	found := false
	index := -1

	for i, num := range numbers {
		if num == target {
			found = true
			index = i
			break // 找到后立即停止搜索
		}
	}

	if found {
		fmt.Printf("找到了！%d 在索引 %d 处\n", target, index)
	} else {
		fmt.Printf("没有找到 %d\n", target)
	}
}

func TestMain8(t *testing.T) {
	scores := []int{85, -1, 92, 0, 78, -1, 95, 88}

	fmt.Println("有效成绩（跳过无效值-1和0）:")
	sum := 0
	count := 0

	for _, score := range scores {
		// 跳过无效成绩
		if score <= 0 {
			continue
		}

		fmt.Println(score)
		sum += score
		count++
	}

	if count > 0 {
		fmt.Printf("\n有效成绩数: %d\n", count)
		fmt.Printf("平均分: %.1f\n", float64(sum)/float64(count))
	}
}
