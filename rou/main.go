package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printNum(ss string) {
	defer wg.Done()
	fmt.Println(ss)
}

func main() {
	wg.Add(5)

	go printNum("打印协程1")
	go printNum("打印协程2")
	go printNum("打印协程3")
	go printNum("打印协程4")
	go printNum("打印协程5")

	wg.Wait()
}
