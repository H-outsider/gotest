package main3

import (
	"fmt"
	"sync"
	"testing"
)

func TestMain1(t *testing.T) {
	// 学生成绩：学生名 -> 科目 -> 分数
	grades := make(map[string]map[string]int)

	// 初始化内层map
	grades["张三"] = make(map[string]int)
	grades["张三"]["语文"] = 85
	grades["张三"]["数学"] = 92

	grades["李四"] = map[string]int{
		"语文": 78,
		"数学": 88,
		"英语": 90,
	}

	// 访问
	fmt.Println("张三的数学:", grades["张三"]["数学"])

	// 遍历
	for name, subjects := range grades {
		fmt.Printf("%s的成绩:\n", name)
		for subject, score := range subjects {
			fmt.Printf("  %s: %d\n", subject, score)
		}
	}
}

func TestMain2(t *testing.T) {
	users := map[int]User{
		1001: {Name: "张三", Age: 25, Email: "zhangsan@example.com"},
		1002: {Name: "李四", Age: 30, Email: "lisi@example.com"},
	}

	// 读取
	user := users[1001]
	fmt.Printf("用户1001: %+v\n", user)

	// 修改整个结构体
	users[1001] = User{Name: "张三", Age: 26, Email: "zhangsan@example.com"}

	// 注意：不能直接修改结构体的字段
	// users[1001].Age = 26  // 编译错误

	// 正确做法：取出、修改、放回
	u := users[1002]
	u.Age = 31
	users[1002] = u

	fmt.Printf("修改后用户1002: %+v\n", users[1002])
}

func TestMain3(t *testing.T) {
	users := map[int]*SimpleUser{
		1: {Name: "张三", Age: 25},
		2: {Name: "李四", Age: 30},
	}

	// 可以直接修改字段
	users[1].Age = 26
	fmt.Printf("用户1: %+v\n", *users[1])

	// 添加新用户
	users[3] = &SimpleUser{Name: "王五", Age: 28}

	for id, user := range users {
		fmt.Printf("ID %d: %s, %d岁\n", id, user.Name, user.Age)
	}
}

func TestMain4(t *testing.T) {
	// 使用map[T]bool实现集合
	set := make(map[string]bool)

	// 添加元素
	set["apple"] = true
	set["banana"] = true
	set["orange"] = true

	// 检查元素是否存在
	if set["apple"] {
		fmt.Println("apple在集合中")
	}

	if !set["grape"] {
		fmt.Println("grape不在集合中")
	}

	// 删除元素
	delete(set, "banana")

	// 遍历集合
	fmt.Println("集合元素:")
	for item := range set {
		fmt.Println(" ", item)
	}

	// 也可以用map[T]struct{}更节省内存
	set2 := make(map[string]struct{})
	set2["a"] = struct{}{}
	if _, ok := set2["a"]; ok {
		fmt.Println("a存在")
	}
}

func TestMain5(t *testing.T) {
	t.Skip("演示用例：多个goroutine同时写入普通map可能触发fatal error: concurrent map writes")

	m := make(map[int]int)
	var wg sync.WaitGroup

	// 多个goroutine同时写入（危险！）
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i // 并发写入，可能panic
		}(i)
	}

	wg.Wait()
	fmt.Println("完成") // 可能到不了这里
}

func TestMain6(t *testing.T) {
	rm := NewRWMap()
	rm.Set("test", 100)

	var wg sync.WaitGroup

	// 多个goroutine并发读取
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if val, ok := rm.Get("test"); ok {
				fmt.Println("读取到:", val)
			}
		}()
	}

	wg.Wait()
}

func TestMain7(t *testing.T) {
	var m sync.Map
	var wg sync.WaitGroup

	// 并发写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(fmt.Sprintf("key%d", i), i*10)
		}(i)
	}

	wg.Wait()

	// 读取
	if val, ok := m.Load("key5"); ok {
		fmt.Println("key5:", val)
	}

	// 遍历
	fmt.Println("所有键值对:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true // 返回false则停止遍历
	})

	// LoadOrStore：不存在则存储
	actual, loaded := m.LoadOrStore("key100", 1000)
	fmt.Printf("LoadOrStore key100: actual=%v, loaded=%v\n", actual, loaded)

	// 删除
	m.Delete("key0")
}
