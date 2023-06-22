package main

import (
	"fmt"
)

func main() {
	// defe.GetTag()

	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // 向通道写入数据
		}
		close(ch) // 关闭通道
	}()
	for {
		value, ok := <-ch
		if !ok {
			break // 通道已关闭，退出循环
		}
		fmt.Println(value) // 从通道读取数据
	}

	// 尝试向关闭的通道写入数据
	ch <- 6 // 这里将导致 panic
}
