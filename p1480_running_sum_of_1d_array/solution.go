package p1480_running_sum_of_1d_array

func runningSum(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	return prefix
}
