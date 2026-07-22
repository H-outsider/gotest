package main8

type Person struct {
	Name string
	Age  int
}

// ByAge 按年龄排序的类型。
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByName 按姓名排序的类型。
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Item struct {
	Name  string
	Value int
	Order int // 原始顺序
}

type Student struct {
	Name  string
	Class string
	Score int
}
