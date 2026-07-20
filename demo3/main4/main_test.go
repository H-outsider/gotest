package main4

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMain1(t *testing.T) {
	// 方式1：指定字段名（推荐）
	p1 := PersonBasic{
		Name:  "张三",
		Age:   25,
		Email: "zhangsan@example.com",
	}
	fmt.Printf("p1: %+v\n", p1)

	// 方式2：按顺序赋值（不推荐，字段顺序变化会出错）
	p2 := PersonBasic{"李四", 30, "lisi@example.com"}
	fmt.Printf("p2: %+v\n", p2)

	// 方式3：部分字段，其他为零值
	p3 := PersonBasic{Name: "王五"}
	fmt.Printf("p3: %+v\n", p3)

	// 方式4：零值结构体
	var p4 PersonBasic
	fmt.Printf("p4: %+v\n", p4)
}

func TestMain2(t *testing.T) {
	p := PersonFields{Name: "张三", Age: 25}

	// 读取字段
	fmt.Println("姓名:", p.Name)
	fmt.Println("年龄:", p.Age)

	// 修改字段
	p.Age = 26
	p.Name = "张三丰"
	fmt.Printf("修改后: %+v\n", p)
}

func TestMain3(t *testing.T) {
	// 创建结构体指针
	p := &PersonPointer{Name: "张三", Age: 25}

	// 通过指针访问字段（自动解引用）
	fmt.Println("姓名:", p.Name) // 等同于 (*p).Name

	// 通过指针修改字段
	p.Age = 26
	fmt.Printf("修改后: %+v\n", *p)

	// 使用new创建
	p2 := new(PersonPointer)
	p2.Name = "李四"
	p2.Age = 30
	fmt.Printf("p2: %+v\n", *p2)
}

func TestMain4(t *testing.T) {
	p1 := NewPerson("张三", 25)
	p2 := NewPersonWithDefaults("李四")

	fmt.Printf("p1: %+v\n", *p1)
	fmt.Printf("p2: %+v\n", *p2)
}

func TestMain5(t *testing.T) {
	// 匿名结构体
	person := struct {
		Name string
		Age  int
	}{
		Name: "临时用户",
		Age:  20,
	}

	fmt.Printf("%+v\n", person)

	// 常用于测试数据
	testCases := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for _, tc := range testCases {
		fmt.Printf("输入: %d, 期望: %d\n", tc.input, tc.expected)
	}
}

func TestMain6(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{1, 3}

	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false
}

func TestMain7(t *testing.T) {
	// 创建商品
	apple := Product{ID: 1, Name: "苹果", Price: 5.5}
	banana := Product{ID: 2, Name: "香蕉", Price: 3.0}

	// 创建订单
	order := Order{
		ID:        1001,
		CreatedAt: time.Now(),
	}

	// 添加商品
	order.AddItem(apple, 3)
	order.AddItem(banana, 5)

	// 输出订单信息
	fmt.Println("===== 订单详情 =====")
	fmt.Println("订单号:", order.ID)
	fmt.Println("创建时间:", order.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("商品列表:")
	for _, item := range order.Items {
		subtotal := item.Product.Price * float64(item.Quantity)
		fmt.Printf("  %s x %d = %.2f元\n",
			item.Product.Name, item.Quantity, subtotal)
	}
	fmt.Printf("总计: %.2f元\n", order.Total())
}

func TestMain8(t *testing.T) {
	user := JSONUser{
		ID:       1,
		Username: "zhangsan",
		Password: "secret123",
		Email:    "",
	}

	data, _ := json.Marshal(user)
	fmt.Println(string(data))
}
