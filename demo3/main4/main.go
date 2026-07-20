package main4

import "time"

// 定义Person结构体
type PersonBasic struct {
	Name  string
	Age   int
	Email string
}

type PersonFields struct {
	Name string
	Age  int
}

type PersonPointer struct {
	Name string
	Age  int
}

type PersonConstructor struct {
	Name string
	Age  int
}

type Point struct {
	X, Y int
}

// 商品
type Product struct {
	ID    int
	Name  string
	Price float64
}

// 订单项
type OrderItem struct {
	Product  Product
	Quantity int
}

// 订单
type Order struct {
	ID        int
	Items     []OrderItem
	CreatedAt time.Time
}

type JSONUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`               // 不序列化
	Email    string `json:"email,omitempty"` // 空则省略
}

// NewPerson 构造函数，返回指针。
func NewPerson(name string, age int) *PersonConstructor {
	return &PersonConstructor{
		Name: name,
		Age:  age,
	}
}

// NewPersonWithDefaults 带默认值的构造函数。
func NewPersonWithDefaults(name string) *PersonConstructor {
	return &PersonConstructor{
		Name: name,
		Age:  18, // 默认年龄
	}
}

// Total 计算订单总价。
func (o *Order) Total() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

// AddItem 添加商品。
func (o *Order) AddItem(product Product, quantity int) {
	o.Items = append(o.Items, OrderItem{
		Product:  product,
		Quantity: quantity,
	})
}
