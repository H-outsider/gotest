package main9

import "fmt"

func checkA() bool {
	fmt.Println("执行了checkA")
	return false
}

func checkB() bool {
	fmt.Println("执行了checkB")
	return true
}
