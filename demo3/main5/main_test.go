package main5

import (
	"fmt"
	"testing"
)

func TestMain1(t *testing.T) {
	// 方式1：取地址
	p1 := PersonPointer{Name: "张三", Age: 25}
	ptr1 := &p1

	// 方式2：直接创建指针
	ptr2 := &PersonPointer{Name: "李四", Age: 30}

	// 方式3：使用new（字段为零值）
	ptr3 := new(PersonPointer)
	ptr3.Name = "王五"
	ptr3.Age = 28

	fmt.Printf("ptr1: %+v\n", *ptr1)
	fmt.Printf("ptr2: %+v\n", *ptr2)
	fmt.Printf("ptr3: %+v\n", *ptr3)
}

func TestMain2(t *testing.T) {
	p := &PersonPointer{Name: "张三", Age: 25}

	// 这两种写法等价
	fmt.Println((*p).Name) // 显式解引用
	fmt.Println(p.Name)    // Go自动解引用（推荐）

	// 修改也可以简化
	p.Age = 26 // 等价于 (*p).Age = 26
	fmt.Printf("修改后: %+v\n", *p)
}

func TestMain3(t *testing.T) {
	p := PersonPointer{Name: "张三", Age: 25}

	birthdayValue(p)
	fmt.Println("值传递后:", p.Age) // 25，没变

	birthdayPointer(&p)
	fmt.Println("指针传递后:", p.Age) // 26，变了
}

func TestMain4(t *testing.T) {
	p := EmbeddedPerson{
		Name: "张三",
		Age:  25,
		Address: Address{
			City:   "北京",
			Street: "长安街",
		},
	}

	// 直接访问嵌入类型的字段（字段提升）
	fmt.Println("城市:", p.City)
	fmt.Println("街道:", p.Street)

	// 也可以通过类型名访问
	fmt.Println("完整地址:", p.Address)
}

func TestMain5(t *testing.T) {
	car := Car{
		Brand:  "Tesla",
		Engine: Engine{Power: 500},
	}

	// 直接调用嵌入类型的方法
	car.Start()
	car.Stop()

	// 也可以通过类型名调用
	car.Engine.Start()
}

func TestMain6(t *testing.T) {
	fc := FlyingCar{
		Model:    "AeroX",
		Flyable:  Flyable{MaxAltitude: 3000},
		Drivable: Drivable{MaxSpeed: 200},
	}

	fmt.Println("型号:", fc.Model)
	fc.Fly()
	fc.Drive()
}

func TestMain7(t *testing.T) {
	c := C{
		A:    A{Name: "来自A"},
		B:    B{Name: "来自B"},
		Name: "来自C",
	}

	// 直接访问是C自己的Name
	fmt.Println("c.Name:", c.Name)

	// 访问嵌入类型的Name需要指定类型
	fmt.Println("c.A.Name:", c.A.Name)
	fmt.Println("c.B.Name:", c.B.Name)
}

func TestMain8(t *testing.T) {
	logger := &Logger{Prefix: "INFO"}

	svc := Service{
		Name:   "UserService",
		Logger: logger,
	}

	svc.Log("服务启动")

	// 多个Service可以共享同一个Logger
	svc2 := Service{
		Name:   "OrderService",
		Logger: logger,
	}
	svc2.Log("服务启动")
}

func TestMain9(t *testing.T) {
	// 创建用户
	user := User{
		Username: "zhangsan",
		Email:    "zhangsan@example.com",
	}
	user.ID = 1
	user.BeforeCreate()

	fmt.Printf("用户: ID=%d, 用户名=%s, 创建时间=%s\n",
		user.ID, user.Username, user.CreatedAt.Format("15:04:05"))

	// 创建文章
	article := Article{
		Title:    "Go语言入门",
		Content:  "Go是一门简洁的语言...",
		AuthorID: user.ID,
	}
	article.ID = 100
	article.BeforeCreate()

	fmt.Printf("文章: ID=%d, 标题=%s, 创建时间=%s\n",
		article.ID, article.Title, article.CreatedAt.Format("15:04:05"))
}
