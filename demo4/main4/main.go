package main4

import (
	"fmt"
	"reflect"
)

type BasicSpeaker interface {
	Speak()
}

type BasicDog struct {
	Name string
}

func (d BasicDog) Speak() {
	fmt.Println(d.Name, "says woof!")
}

type BasicCat struct {
	Name string
}

func (c BasicCat) Speak() {
	fmt.Println(c.Name, "says meow!")
}

type NilSpeaker interface {
	Speak()
}

type NilDog struct {
	Name string
}

func (d *NilDog) Speak() {
	if d == nil {
		fmt.Println("nil dog")
		return
	}
	fmt.Println(d.Name, "says woof!")
}

type ComparableSpeaker interface {
	Speak()
}

type ComparableDog struct {
	Name string
}

func (d ComparableDog) Speak() {}

type Container interface {
	Get() interface{}
}

type SliceContainer struct {
	Data []int
}

func (s SliceContainer) Get() interface{} {
	return s.Data
}

type Animal interface {
	Speak() string
}

type InspectDog struct {
	Name string
	Age  int
}

func (d InspectDog) Speak() string {
	return "Woof!"
}

type InspectCat struct {
	Name string
}

func (c *InspectCat) Speak() string {
	return "Meow!"
}

func inspectInterface(a Animal) {
	if a == nil {
		fmt.Println("接口是nil")
		return
	}

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	fmt.Printf("动态类型: %v\n", t)
	fmt.Printf("动态值: %v\n", v)
	fmt.Printf("是否指针: %v\n", t.Kind() == reflect.Ptr)

	// 如果是指针，检查是否为nil
	if t.Kind() == reflect.Ptr {
		fmt.Printf("指针是否为nil: %v\n", v.IsNil())
	}
	fmt.Println("---")
}
