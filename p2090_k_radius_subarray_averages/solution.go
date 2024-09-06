package p2090_k_radius_subarray_averages

func getAverages(nums []int, k int) []int {
	n := len(nums)
	prefix := make([]int, n+1)
	prefix[0] = 0
	avg := make([]int, n)

	for i := 1; i < n+1; i++ {
		prefix[i] += prefix[i-1] + nums[i-1]
		avg[i-1] = -1
	}

	for i := k; i+k < n; i++ {
		left := prefix[i] - prefix[i-k]
		right := prefix[i+k+1] - prefix[i]
		avg[i] = (left + right) / (k*2 + 1)
	}

	return avg
}
