package main5

import (
	"fmt"
	"time"
)

type PersonPointer struct {
	Name string
	Age  int
}

// birthdayValue 值传递：不会修改原结构体。
func birthdayValue(p PersonPointer) {
	p.Age++
}

// birthdayPointer 指针传递：会修改原结构体。
func birthdayPointer(p *PersonPointer) {
	p.Age++
}

// 基础类型
type Address struct {
	City   string
	Street string
}

// EmbeddedPerson 嵌入Address。
type EmbeddedPerson struct {
	Name string
	Age  int
	Address
}

type Engine struct {
	Power int
}

func (e *Engine) Start() {
	fmt.Println("引擎启动，功率:", e.Power, "马力")
}

func (e *Engine) Stop() {
	fmt.Println("引擎关闭")
}

type Car struct {
	Brand string
	Engine
}

type Flyable struct {
	MaxAltitude int
}

func (f *Flyable) Fly() {
	fmt.Println("飞行中，最大高度:", f.MaxAltitude, "米")
}

type Drivable struct {
	MaxSpeed int
}

func (d *Drivable) Drive() {
	fmt.Println("行驶中，最大速度:", d.MaxSpeed, "公里/时")
}

// FlyingCar 嵌入多个类型。
type FlyingCar struct {
	Model string
	Flyable
	Drivable
}

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
	Name string
}

type Logger struct {
	Prefix string
}

func (l *Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

type Service struct {
	Name string
	*Logger
}

// BaseModel 基础模型，包含公共字段。
type BaseModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseModel) BeforeCreate() {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *BaseModel) BeforeUpdate() {
	b.UpdatedAt = time.Now()
}

// User 用户模型。
type User struct {
	BaseModel
	Username string
	Email    string
}

// Article 文章模型。
type Article struct {
	BaseModel
	Title    string
	Content  string
	AuthorID int
}
