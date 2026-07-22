package main3

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Kitty"}

	MakeSound(dog)
	MakeSound(cat)
}

func TestMain2(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	PrintArea(rect)
	PrintArea(circle)

	shapes := []Shape{rect, circle}
	fmt.Printf("总面积: %.2f\n", TotalArea(shapes))
}

func TestMain3(t *testing.T) {
	db := NewDatabase("mysql", "localhost:3306")
	db.Connect()
	results, _ := db.Query("SELECT * FROM users")
	fmt.Println("查询结果:", results)
	db.Close()
}

func TestMain4(t *testing.T) {
	file := &File{Name: "test.txt"}

	// File同时实现了Reader、Writer和ReadWriter
	var r Reader = file
	var w Writer = file
	var rw ReadWriter = file

	w.Write("Hello, World!")
	fmt.Println("Read:", r.Read())

	rw.Write("Updated content")
	fmt.Println("ReadWriter Read:", rw.Read())
}

func TestMain5(t *testing.T) {
	card := &CreditCard{
		CardNumber: "6222021234567890",
		Balance:    1000,
	}

	alipay := &Alipay{
		Account: "user@example.com",
		Balance: 500,
	}

	order1 := &Order{ID: "ORD001", Amount: 99.9}
	order2 := &Order{ID: "ORD002", Amount: 199.9}

	// 不同订单使用不同支付方式
	_ = order1.Checkout(card)
	_ = order2.Checkout(alipay)

	fmt.Println("\n余额查询:")
	fmt.Printf("信用卡余额: %.2f\n", card.Balance)
	fmt.Printf("支付宝余额: %.2f\n", alipay.Balance)
}
