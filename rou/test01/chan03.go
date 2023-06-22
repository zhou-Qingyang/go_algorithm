package test01

import (
	"fmt"
	"time"
)

func loop(ch chan int) {
	i := <-ch
	fmt.Println("接收数据", i)
}

func PrintNum() {
	ch1 := make(chan int)
	go loop(ch1)
	time.Sleep(time.Duration(3) * time.Second) //我在这睡着了 主线程不动 loop一直在等待传值 所以不会执行接受数据
	ch1 <- 1
	time.Sleep(time.Duration(3) * time.Second)
}
