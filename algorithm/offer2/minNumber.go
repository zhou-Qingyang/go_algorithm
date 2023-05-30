package offer02

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	stringNum := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		stringNum[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(stringNum, func(i, j int) bool {
		return stringNum[i]+stringNum[j] <= stringNum[i]+stringNum[j]
	})
	res := ""
	for _, re := range stringNum {
		res += re
	}
	return res
}
