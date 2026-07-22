package main2

import "fmt"

type MethodPerson struct {
	Name string
}

func (p MethodPerson) SayHello() {
	fmt.Println("Hello, I'm", p.Name)
}

func (p *MethodPerson) SetName(name string) {
	p.Name = name
}

type PointerGreeter interface {
	Greet()
}

type PointerGreeterPerson struct {
	Name string
}

// Greet 指针接收者方法。
func (p *PointerGreeterPerson) Greet() {
	fmt.Println("Hello, I'm", p.Name)
}

type ValueGreeter interface {
	Greet()
}

type ValueGreeterPerson struct {
	Name string
}

// Greet 值接收者方法。
func (p ValueGreeterPerson) Greet() {
	fmt.Println("Hello, I'm", p.Name)
}

type Animal interface {
	Speak()
	Move()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "says: Woof!")
}

func (d *Dog) Move() {
	fmt.Println(d.Name, "is running")
}

// Counter 定义接口。
type Counter interface {
	Increment()
	Decrement()
	Value() int
}

// SimpleCounter 实现一。
type SimpleCounter struct {
	count int
}

func (c *SimpleCounter) Increment() {
	c.count++
}

func (c *SimpleCounter) Decrement() {
	c.count--
}

func (c *SimpleCounter) Value() int {
	return c.count
}

// UseCounter 使用接口的函数。
func UseCounter(c Counter) {
	c.Increment()
	c.Increment()
	c.Increment()
	c.Decrement()
	fmt.Println("最终值:", c.Value())
}
