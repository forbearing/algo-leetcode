package array

func SortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	l, r, i := 0, len(nums)-1, len(nums)-1

	for l <= r {
		if nums[l]*nums[l] < nums[r]*nums[r] {
			result[i] = nums[r] * nums[r]
			r--
		} else {
			result[i] = nums[l] * nums[l]
			l++
		}
		i--
	}
	return result
}
