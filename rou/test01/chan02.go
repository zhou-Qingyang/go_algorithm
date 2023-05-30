package test01

import "fmt"

// 打印 1 的函数
func printOne(c1, c3 chan int, done chan bool) {
	for i := 0; i < 100; i++ {
		<-c3
		fmt.Println(1)

		// 在通道已满时跳过写操作
		select {
		case c1 <- 1:
		default:
		}
	}
	done <- true
}

// 打印 2 的函数
func printTwo(c1, c2 chan int, done chan bool) {
	for i := 0; i < 100; i++ {
		<-c1
		fmt.Println(2)

		// 在通道已满时跳过写操作
		select {
		case c2 <- 1:
		default:
		}
	}
	done <- true
}

// 打印 3 的函数
func printThree(c2, c3 chan int, done chan bool) {
	for i := 0; i < 100; i++ {
		<-c2
		fmt.Println(3)

		// 在通道已满时跳过写操作
		select {
		case c3 <- 1:
		default:
		}
	}
	done <- true
}

func PrintNum3() {
	c1 := make(chan int, 1) // 将 c1 初始化为带缓冲的通道，避免协程死锁
	c2 := make(chan int, 1) // 将 c2 初始化为带缓冲的通道，避免协程死锁
	c3 := make(chan int, 1) // 将 c3 初始化为带缓冲的通道，避免协程死锁

	// 用于标识协程是否已经完成工作
	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)

	// 初始化 c3
	go func() {
		c3 <- 1
	}()

	// 启动协程
	for i := 0; i < 100; i++ {
		go printOne(c1, c3, done1)
		go printTwo(c1, c2, done2)
		go printThree(c2, c3, done3)
	}

	// 等待所有协程完成工作
	for i := 0; i < 300; i++ {
		select {
		case <-done1:
			if i == 299 {
				close(c1)
			}
		case <-done2:
			if i == 299 {
				close(c2)
			}
		case <-done3:
			if i == 299 {
				close(c3)
			}
		}
	}
}
