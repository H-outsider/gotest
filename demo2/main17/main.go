package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Config struct {
	Host     string
	Port     int
	Debug    bool
	MaxConns int
}

// 将变量值翻倍
func double(n *int) {
	*n = *n * 2
}

// 交换两个变量的值
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 同时修改两个变量
func addBoth(x, y *int, delta int) {
	*x += delta
	*y += delta
}

// 同时计算商和余数
func divmod(a, b int, quotient, remainder *int) {
	*quotient = a / b
	*remainder = a % b
}

// 值传递：不会修改原结构体
func birthdayValue(p Person) {
	p.Age++
	fmt.Println("函数内（值传递）:", p.Age)
}

// 指针传递：会修改原结构体
func birthdayPointer(p *Person) {
	p.Age++ // Go允许直接用.访问，相当于(*p).Age++
	fmt.Println("函数内（指针传递）:", p.Age)
}

// 计数器操作
func increment(count *int) {
	*count++
}

func decrement(count *int) {
	*count--
}

func reset(count *int) {
	*count = 0
}

// 使用指针参数设置默认值
func setDefaults(c *Config) {
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.MaxConns == 0 {
		c.MaxConns = 100
	}
}
