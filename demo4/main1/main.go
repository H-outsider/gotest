package main1

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Width  float64
	Height float64
}

// Area 值接收者方法。
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Counter struct {
	Value int
}

// IncrementWrong 值接收者：不能修改原值。
func (c Counter) IncrementWrong() {
	c.Value++ // 修改的是副本
}

// Increment 指针接收者：可以修改原值。
func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) Add(n int) {
	c.Value += n
}

type Point struct {
	X, Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// MyInt 给int定义新类型。
type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}

func (m MyInt) Double() MyInt {
	return m * 2
}

// IntSlice 给切片定义新类型。
type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

type Account struct {
	Owner   string
	Balance float64
}

// NewAccount 创建账户。
func NewAccount(owner string, initial float64) (*Account, error) {
	if initial < 0 {
		return nil, errors.New("初始金额不能为负")
	}
	return &Account{Owner: owner, Balance: initial}, nil
}

// Deposit 存款。
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("存款金额必须为正数")
	}
	a.Balance += amount
	return nil
}

// Withdraw 取款。
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("取款金额必须为正数")
	}
	if amount > a.Balance {
		return errors.New("余额不足")
	}
	a.Balance -= amount
	return nil
}

// GetBalance 查询余额。
func (a Account) GetBalance() float64 {
	return a.Balance
}

// Transfer 转账。
func (a *Account) Transfer(to *Account, amount float64) error {
	if err := a.Withdraw(amount); err != nil {
		return fmt.Errorf("转账失败: %w", err)
	}
	if err := to.Deposit(amount); err != nil {
		return fmt.Errorf("转账失败: %w", err)
	}
	return nil
}

// String 账户信息。
func (a Account) String() string {
	return fmt.Sprintf("账户[%s] 余额: %.2f元", a.Owner, a.Balance)
}
