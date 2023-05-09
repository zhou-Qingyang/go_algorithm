package day01

//1
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		m[v] = i
	}
	return nil

}
