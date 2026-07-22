package main2

import "testing"

func TestMain1(t *testing.T) {
	// 值类型可以调用两种方法（直接调用）
	p1 := MethodPerson{Name: "Alice"}
	p1.SayHello()     // 值接收者，正常
	p1.SetName("Bob") // 指针接收者，Go自动取地址
	p1.SayHello()

	// 指针类型也可以调用两种方法
	p2 := &MethodPerson{Name: "Charlie"}
	p2.SayHello() // 值接收者，Go自动解引用
	p2.SetName("David")
	p2.SayHello()
}

func TestMain2(t *testing.T) {
	var g PointerGreeter

	// 指针可以赋值给接口（*Person的方法集包含Greet）
	g = &PointerGreeterPerson{Name: "Alice"}
	g.Greet()

	// 值不能赋值给接口（Person的方法集不包含Greet）
	// g = PointerGreeterPerson{Name: "Bob"}  // 编译错误！
}

func TestMain3(t *testing.T) {
	var g ValueGreeter

	// 值可以赋值给接口
	g = ValueGreeterPerson{Name: "Alice"}
	g.Greet()

	// 指针也可以赋值给接口
	g = &ValueGreeterPerson{Name: "Bob"}
	g.Greet()
}

func TestMain4(t *testing.T) {
	// Dog值的方法集只有Speak，没有Move
	// 所以Dog值不能实现Animal接口

	// *Dog的方法集有Speak和Move
	// 所以*Dog可以实现Animal接口
	var a Animal = &Dog{Name: "Buddy"}
	a.Speak()
	a.Move()

	// var a2 Animal = Dog{Name: "Max"}  // 编译错误
}

func TestMain5(t *testing.T) {
	// 必须传指针，因为Increment和Decrement是指针接收者
	counter := &SimpleCounter{}
	UseCounter(counter)

	// 错误示例（会编译失败）
	// counter2 := SimpleCounter{}
	// UseCounter(counter2)  // SimpleCounter没有实现Counter接口
}
