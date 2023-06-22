package defe

import (
	"fmt"
	"time"
)

func Proc() {
	panic("这个报错了")
}

func Sloution(seconds int) {
	go func() {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("出错了", err)
			}
		}()
		Proc()

	}()
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("开始执行代码")
}
