package main3

import (
	"errors"
	"fmt"
	"math"
)

// Speaker 定义接口。
type Speaker interface {
	Speak() string
}

// Dog 实现Speaker。
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

// Cat 实现Speaker。
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

// MakeSound 使用接口的函数。
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// PrintArea 接受任何实现了Shape接口的类型。
func PrintArea(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
}

// TotalArea 接受Shape切片。
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

type Database interface {
	Connect() error
	Query(sql string) ([]string, error)
	Close()
}

type MySQL struct {
	Host string
}

func (m *MySQL) Connect() error {
	fmt.Println("连接到MySQL:", m.Host)
	return nil
}

func (m *MySQL) Query(sql string) ([]string, error) {
	return []string{"result1", "result2"}, nil
}

func (m *MySQL) Close() {
	fmt.Println("关闭MySQL连接")
}

type PostgreSQL struct {
	Host string
}

func (p *PostgreSQL) Connect() error {
	fmt.Println("连接到PostgreSQL:", p.Host)
	return nil
}

func (p *PostgreSQL) Query(sql string) ([]string, error) {
	return []string{"pg_result1"}, nil
}

func (p *PostgreSQL) Close() {
	fmt.Println("关闭PostgreSQL连接")
}

// NewDatabase 工厂函数返回接口。
func NewDatabase(dbType, host string) Database {
	switch dbType {
	case "mysql":
		return &MySQL{Host: host}
	case "postgres":
		return &PostgreSQL{Host: host}
	default:
		return nil
	}
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Name    string
	Content string
}

func (f *File) Read() string {
	return f.Content
}

func (f *File) Write(data string) {
	f.Content = data
}

// PaymentMethod 支付接口。
type PaymentMethod interface {
	Pay(amount float64) error
	Refund(amount float64) error
	GetName() string
}

// CreditCard 信用卡支付。
type CreditCard struct {
	CardNumber string
	Balance    float64
}

func (c *CreditCard) Pay(amount float64) error {
	if amount > c.Balance {
		return errors.New("余额不足")
	}
	c.Balance -= amount
	fmt.Printf("信用卡支付 %.2f 元成功\n", amount)
	return nil
}

func (c *CreditCard) Refund(amount float64) error {
	c.Balance += amount
	fmt.Printf("信用卡退款 %.2f 元成功\n", amount)
	return nil
}

func (c *CreditCard) GetName() string {
	return "信用卡 " + c.CardNumber[len(c.CardNumber)-4:]
}

// Alipay 支付宝支付。
type Alipay struct {
	Account string
	Balance float64
}

func (a *Alipay) Pay(amount float64) error {
	if amount > a.Balance {
		return errors.New("支付宝余额不足")
	}
	a.Balance -= amount
	fmt.Printf("支付宝支付 %.2f 元成功\n", amount)
	return nil
}

func (a *Alipay) Refund(amount float64) error {
	a.Balance += amount
	fmt.Printf("支付宝退款 %.2f 元成功\n", amount)
	return nil
}

func (a *Alipay) GetName() string {
	return "支付宝 " + a.Account
}

// Order 订单处理。
type Order struct {
	ID     string
	Amount float64
}

func (o *Order) Checkout(payment PaymentMethod) error {
	fmt.Printf("订单 %s，金额 %.2f，使用 %s 支付\n",
		o.ID, o.Amount, payment.GetName())
	return payment.Pay(o.Amount)
}
