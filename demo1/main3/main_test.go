package main3

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	fmt.Println("Hello World!")
}

func TestMain2(t *testing.T) {
	fmt.Println("第一行")
	fmt.Println("第二行")
	fmt.Println("第三行")
}

func TestMain3(t *testing.T) {
	fmt.Print("Hello,")
	fmt.Print("World!")
	fmt.Println()
}

func TestMain4(t *testing.T) {
	name := "张三"
	age := 25
	fmt.Printf("姓名：%s，年龄：%d", name, age)
}

func TestMain5(t *testing.T) {
	message := "Hello,Go 1.25!"
	fmt.Println(message)
}
