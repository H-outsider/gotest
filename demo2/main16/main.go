package main16

import "fmt"

func addOneValue(n int) {
	n = n + 1
	fmt.Println("函数内 n =", n)
}

// 指针传递：函数收到的是地址，可以修改原变量
func addOnePointer(n *int) {
	*n = *n + 1
	fmt.Println("函数内 *n =", *n)
}

// 使用指针交换两个变量的值
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
