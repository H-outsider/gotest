package main16

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	a := 42

	fmt.Println("变量a的值：", a)
	fmt.Println("变量a的地址：", &a)

	//声明一个指针变量
	var p *int
	p = &a //让p指向a

	fmt.Println("变量p的值（a的地址）：", p)
	fmt.Println("变量p的指针：", *p)
}

func TestMain2(t *testing.T) {
	a := 10
	p := &a // p指向a

	fmt.Println("修改前: a =", a)

	*p = 20 // 通过指针修改a的值

	fmt.Println("修改后: a =", a)
}

func TestMain3(t *testing.T) {
	// new(T) 分配内存，返回指向零值的指针
	p := new(int)

	fmt.Println("new创建的指针:", p)
	fmt.Println("指向的值（零值）:", *p)

	*p = 100
	fmt.Println("修改后:", *p)
}

func TestMain4(t *testing.T) {
	var p *int

	fmt.Println("未初始化的指针:", p)
	fmt.Println("是否为nil:", p == nil)

	// *p = 10  // 危险！对nil指针解引用会panic

	// 使用前应该检查
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("指针为nil，不能使用")
	}

}

func TestMain5(t *testing.T) {
	a := 10

	fmt.Println("=== 值传递 ===")
	fmt.Println("调用前 a =", a)
	addOneValue(a)
	fmt.Println("调用后 a =", a) // a没变

	fmt.Println("\n=== 指针传递 ===")
	fmt.Println("调用前 a =", a)
	addOnePointer(&a)
	fmt.Println("调用后 a =", a) // a变了
}

func TestMain6(t *testing.T) {
	x, y := 10, 20

	fmt.Println("交换前: x =", x, ", y =", y)
	swap(&x, &y)
	fmt.Println("交换后: x =", x, ", y =", y)
}
