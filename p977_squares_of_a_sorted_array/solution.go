package p977_squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			res[i] = leftSquare
			left++
		} else {
			res[i] = rightSquare
			right--
		}

	}

	return res
}
