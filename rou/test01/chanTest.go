package test01

// // 打印 1 的函数
// func printOne(c1, c3 chan int, done chan bool) {
// 	for i := 0; i < 5; i++ {
// 		<-c3
// 		fmt.Println(1)
// 		c1 <- 1
// 	}
// 	done <- true
// }

// // 打印 2 的函数
// func printTwo(c1, c2 chan int, done chan bool) {
// 	for i := 0; i < 5; i++ {
// 		<-c1
// 		fmt.Println(2)
// 		c2 <- 1
// 	}
// 	done <- true
// }

// // 打印 3 的函数
// func printThree(c2, c3 chan int, done chan bool) {
// 	for i := 0; i < 5; i++ {
// 		<-c2
// 		fmt.Println(3)
// 		c3 <- 1
// 	}
// 	done <- true
// }

// func PrintNum2() {
// 	c1 := make(chan int)
// 	c2 := make(chan int)
// 	c3 := make(chan int)

// 	// 用于标识协程是否已经完成工作
// 	done1 := make(chan bool)
// 	done2 := make(chan bool)
// 	done3 := make(chan bool)

// 	// 初始化 c3
// 	go func() {
// 		c3 <- 1
// 	}()

// 	// 启动协程
// 	for i := 0; i < 5; i++ {
// 		go printOne(c1, c3, done1)
// 		go printTwo(c1, c2, done2)
// 		go printThree(c2, c3, done3)
// 	}

// 	// 等待所有协程完成工作
// 	for i := 0; i < 15; i++ {
// 		select {
// 		case <-done1:
// 			if i == 14 {
// 				close(c1)
// 			}
// 		case <-done2:
// 			if i == 14 {
// 				close(c2)
// 			}
// 		case <-done3:
// 			if i == 14 {
// 				close(c3)
// 			}
// 		}
// 	}
// }
