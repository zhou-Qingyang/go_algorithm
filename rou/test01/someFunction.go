package test01

import (
	"context"
	"fmt"
	"time"
)

func SomeFunction() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	time.Sleep(5 * time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		//执行 超时的操作
		fmt.Println("1")
	default:
		// 操作完成继续后面的操作
		fmt.Println("2")
	}
}
