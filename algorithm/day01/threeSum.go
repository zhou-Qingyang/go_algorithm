package day01

import "sort"

//15
func TreeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] { // 如果当前元素与前一个元素相同，则跳过此次循环，直接进入下一轮循环
			continue
		}
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] { // 如果 left 指向的元素与下一个元素相同，则向右移动 left 指针
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 如果 right 指向的元素与上一个元素相同，则向左移动 right 指针
					right--
				}
				left++  // 向右移动 left 指针
				right-- // 向左移动 right 指针
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return result // 返回最终结果切片
}
