package p1413_minimum_value_to_get_positive_step_by_step_sum

func minStartValue(nums []int) int {
	sum := 0
	minSum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < minSum {
			minSum = sum
		}
	}

	return 1 - minSum
}
