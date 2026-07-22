package main1

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println("面积:", rect.Area())
	fmt.Println("周长:", rect.Perimeter())
}

func TestMain2(t *testing.T) {
	counter := Counter{Value: 0}

	counter.IncrementWrong()
	fmt.Println("IncrementWrong后:", counter.Value) // 0

	counter.Increment()
	fmt.Println("Increment后:", counter.Value) // 1

	counter.Add(10)
	fmt.Println("Add(10)后:", counter.Value) // 11
}

func TestMain3(t *testing.T) {
	// 值调用指针方法（Go自动取地址）
	p1 := Point{X: 1, Y: 2}
	p1.Move(3, 4) // 等价于 (&p1).Move(3, 4)
	fmt.Println("p1:", p1.String())

	// 指针调用值方法（Go自动解引用）
	p2 := &Point{X: 10, Y: 20}
	fmt.Println("p2:", p2.String()) // 等价于 (*p2).String()
}

func TestMain4(t *testing.T) {
	var n MyInt = 5
	fmt.Println("是否为正:", n.IsPositive())
	fmt.Println("翻倍:", n.Double())

	nums := IntSlice{1, 2, 3, 4, 5}
	fmt.Println("切片求和:", nums.Sum())
}

func TestMain5(t *testing.T) {
	// 创建账户
	alice, _ := NewAccount("Alice", 1000)
	bob, _ := NewAccount("Bob", 500)

	fmt.Println("初始状态:")
	fmt.Println(alice)
	fmt.Println(bob)

	// 存款
	_ = alice.Deposit(200)
	fmt.Println("\nAlice存款200后:", alice.GetBalance())

	// 取款
	err := bob.Withdraw(600)
	if err != nil {
		fmt.Println("Bob取款失败:", err)
	}

	// 转账
	err = alice.Transfer(bob, 300)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("\n转账300后:")
		fmt.Println(alice)
		fmt.Println(bob)
	}
}
