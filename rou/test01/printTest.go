package test01

import (
	"fmt"
	"sync"
)

func PrintNum(id int, wg *sync.WaitGroup) {
	//这种 并没真正意义上 实现同步
	fmt.Printf("Worker %d starting\n", id)
	// // 模拟工作
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d working on task %d\n", id, i)
	}
	fmt.Printf("Worker %d done\n", id)

	wg.Done() // 表示当前 goroutine 已完成
}

func PrintMain() {
	var wg sync.WaitGroup
	numTask := 3

	for i := 0; i < numTask; i++ {
		wg.Add(1)
		go PrintNum(i+1, &wg)
	}
	wg.Wait()
	fmt.Println("All workers done!")
}
