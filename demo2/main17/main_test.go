package main

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	num := 10
	fmt.Println("翻倍前:", num)

	double(&num) // 传入地址
	fmt.Println("翻倍后:", num)
}

func TestMain2(t *testing.T) {
	a, b := 10, 20
	fmt.Println("交换前: a =", a, ", b =", b)
	swap(&a, &b)
	fmt.Println("交换后: a =", a, ", b =", b)

	fmt.Println()

	x, y := 100, 200
	fmt.Println("加之前: x =", x, ", y =", y)
	addBoth(&x, &y, 50)
	fmt.Println("加之后: x =", x, ", y =", y)
}

func TestMain3(t *testing.T) {
	var q, r int
	divmod(17, 5, &q, &r)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", q, r)
}

func TestMain4(t *testing.T) {
	person := Person{Name: "张三", Age: 25}

	fmt.Println("初始年龄:", person.Age)

	birthdayValue(person)
	fmt.Println("值传递后:", person.Age) // 没变

	birthdayPointer(&person)
	fmt.Println("指针传递后:", person.Age) // 变了
}

func TestMain5(t *testing.T) {
	counter := 0

	fmt.Println("初始值:", counter)

	increment(&counter)
	increment(&counter)
	increment(&counter)
	fmt.Println("加3次后:", counter)

	decrement(&counter)
	fmt.Println("减1次后:", counter)

	reset(&counter)
	fmt.Println("重置后:", counter)
}

func TestMain6(t *testing.T) {
	cfg := Config{
		Debug: true,
	}

	fmt.Println("设置默认值前:")
	fmt.Printf("Host: %s, Port: %d, Debug: %t, MaxConns: %d\n",
		cfg.Host, cfg.Port, cfg.Debug, cfg.MaxConns)

	setDefaults(&cfg)

	fmt.Println("\n设置默认值后:")
	fmt.Printf("Host: %s, Port: %d, Debug: %t, MaxConns: %d\n",
		cfg.Host, cfg.Port, cfg.Debug, cfg.MaxConns)
}
