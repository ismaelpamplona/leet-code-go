package p2270_number_of_ways_to_split_array

func waysToSplitArray1(nums []int) int {
	prefix := make([]int, len(nums)+1)
	count := 0
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	for i := 0; i < len(nums)-1; i++ {
		left := prefix[i+1] - prefix[i]
		right := prefix[len(nums)] - prefix[i+1]

		if left >= right {
			count++
		}
	}

	return count
}

func waysToSplitArray2(nums []int) int {
	totalSum := nums[0]
	count := 0

	for i := 1; i < len(nums); i++ {
		totalSum += nums[i]
	}

	left := 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		right := totalSum - left

		if left >= right {
			count++
		}
	}

	return count
}
