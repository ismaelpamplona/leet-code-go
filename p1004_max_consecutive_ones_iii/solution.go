package p1004_max_consecutive_ones_iii

func longestOnes(nums []int, k int) int {
	flips, left, max := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			flips++
		}

		for flips > k {
			if nums[left] == 0 {
				flips--
			}
			left++
		}

		curWindowSize := right - left + 1

		if max < curWindowSize {
			max = curWindowSize
		}
	}

	return max
}
