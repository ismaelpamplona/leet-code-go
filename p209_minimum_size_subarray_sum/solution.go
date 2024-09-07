package p209_minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, sum, minLen := 0, 0, n+1
	for right := 0; right < n; right++ {
		sum += nums[right]

		for sum >= target {
			window := right - left + 1
			if minLen > window {
				minLen = window
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == n+1 {
		minLen = 0
	}
	return minLen
}
