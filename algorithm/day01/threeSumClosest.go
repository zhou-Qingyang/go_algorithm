package day01

import (
	"math"
	"sort"
)

//16.最接近的三数之和
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := math.MaxInt32 // 最接近的差值
	res := 0              // 最接近的三个数的和
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if abs(sum-target) < diff {
				diff = abs(sum - target)
				res = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
