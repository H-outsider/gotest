package main10

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex
)

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // 确保文件被关闭

	// 读取文件内容...
	buf := make([]byte, 100)
	n, err := file.Read(buf)
	if err != nil {
		return err // 即使这里返回，文件也会被关闭
	}

	fmt.Printf("读取了 %d 字节\n", n)
	return nil
}

func increment() {
	mutex.Lock()
	defer mutex.Unlock() // 确保解锁

	counter++
	// 即使中间有panic或return，锁也会被释放
}

func example1() int {
	x := 10
	defer func() {
		x = 20 // 不影响返回值
	}()
	return x // 返回10
}

func example2() (x int) {
	x = 10
	defer func() {
		x = 20 // 修改命名返回值
	}()
	return x // 返回20
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s 耗时 %s\n", name, elapsed)
}

func slowOperation() {
	defer timeTrack(time.Now(), "slowOperation")

	// 模拟耗时操作
	time.Sleep(100 * time.Millisecond)
	fmt.Println("操作完成")
}

type DB struct{}

func (db *DB) Begin() *Tx {
	fmt.Println("开始事务")
	return &Tx{}
}

type Tx struct {
	committed bool
}

func (tx *Tx) Commit() {
	tx.committed = true
	fmt.Println("提交事务")
}

func (tx *Tx) Rollback() {
	if !tx.committed {
		fmt.Println("回滚事务")
	}
}

func doTransaction(db *DB, shouldFail bool) error {
	tx := db.Begin()
	defer tx.Rollback() // 如果没有commit，就rollback

	fmt.Println("执行操作...")

	if shouldFail {
		return fmt.Errorf("操作失败")
	}

	tx.Commit()
	return nil
}
