package test01

import "fmt"

func init() {
	fmt.Println("test01 测试执行了")
}

// 打印 1 的函数
func printOne(c1, c3 chan int, done chan bool) {
	for i := 0; i < 15; i++ {
		<-c3 // 阻塞当前状态
		fmt.Printf("第一个协程:1\n")
		// 在通道已满时跳过写操作
		c1 <- 1
	}
	done <- true
}

// 打印 2 的函数
func printTwo(c1, c2 chan int) {
	for i := 0; i < 15; i++ {
		<-c1
		fmt.Printf("第二个协程:2\n")
		// 在通道已满时跳过写操作
		c2 <- 1
	}

}

// 打印 3 的函数
func printThree(c2, c3 chan int) {
	for i := 0; i < 15; i++ {
		<-c2
		fmt.Printf("第三个协程:3\n")
		c3 <- 1
	}
}

func PrintNum3() {
	c1 := make(chan int, 1) // 将 c1 初始化为带缓冲的通道，避免协程死锁
	c2 := make(chan int, 1) // 将 c2 初始化为带缓冲的通道，避免协程死锁
	c3 := make(chan int, 1) // 将 c3 初始化为带缓冲的通道，避免协程死锁

	done1 := make(chan bool)

	// 初始化 开启线程
	c3 <- 1
	go printOne(c1, c3, done1)
	go printTwo(c1, c2)
	go printThree(c2, c3)

	<-done1 // 这里要先执行
	// 因为子协程更慢
	// 而且 done1这个 是无缓冲的
	// 所以 数据发送前 必须先是接收

}
