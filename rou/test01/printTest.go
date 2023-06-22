package test01

import (
	"fmt"
	"sync"
)

func Printworker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	// 模拟一些耗时的工作
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d working...\n", id)
	}
	fmt.Printf("Worker %d done\n", id)
	// 通知 WaitGroup 完成了一项任务
	wg.Done()
}

func Printmain() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		// 需要等待 3 个 goroutine 完成任务
		wg.Add(1)
		// 启动一个 goroutine 执行 worker 函数，传入 WaitGroup 的地址
		go Printworker(i, &wg)
	}
	fmt.Println("Waiting for all workers to finish...")
	// 阻塞主 goroutine，等待所有 WaitGroup 中的任务完成
	wg.Wait()
	fmt.Println("All workers finished")
}
