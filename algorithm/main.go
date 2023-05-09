package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "0103293034"
	result := [5]string{}

	for i, char := range numStr {
		// 每隔 2 个字符更新数字下标
		if i%2 == 0 {
			//判断第一个数字 是不是0
			num, _ := strconv.Atoi(string(char))
			var res string
			if num == 0 {
				res = numStr[i+1 : i+2]
			} else {
				res = numStr[i : i+2]
			}
			result[i/2] = res
		}
	}
	fmt.Printf("数组:%s", result)
}
