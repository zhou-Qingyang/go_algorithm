package offer

//12
func minArray(numbers []int) int {

	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) / 2
		if numbers[right] > numbers[mid] {
			right = mid - 1
		} else if numbers[right] < numbers[mid] {
			left = mid + 1
		} else if numbers[right] == numbers[mid] {
			right--
		}
	}
	return numbers[left]

}
