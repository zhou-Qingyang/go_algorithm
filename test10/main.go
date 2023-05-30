package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// 读取文件内容
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将文本数据分割成行
	lines := strings.Split(string(data), "\n")

	// 遍历每一行数据，并按需解析其中的字段
	for _, line := range lines {
		fields := strings.Split(line, "\t")
		redStr := fields[2]
		resStr := []string{}
		for i, char := range redStr {
			// 每隔 2 个字符更新数字下标
			if i%2 == 0 {
				//判断第一个数字 是不是0
				num, _ := strconv.Atoi(string(char))
				var res string
				if num == 0 {
					res = redStr[i+1 : i+2]
				} else {
					res = redStr[i : i+2]
				}
				resStr = append(resStr, res)
			}
		}
		blueStr := fields[3]
		for i, char := range blueStr {
			// 每隔 2 个字符更新数字下标
			if i%2 == 0 {
				//判断第一个数字 是不是0
				num, _ := strconv.Atoi(string(char))
				var res string
				if num == 0 {
					res = blueStr[i+1 : i+2]
				} else {
					res = blueStr[i : i+2]
				}
				resStr = append(resStr, res)
			}
		}

		fmt.Printf("Y%s=%2s\n", fields[0], resStr)

	}
	fmt.Println("=========================")
}
