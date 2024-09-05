package p713_subarray_product_less_than_k

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	ans, left, cur := 0, 0, 1

	for right := 0; right < len(nums); right++ {
		cur *= nums[right]
		for cur >= k {
			cur /= nums[left]
			left++
		}

		ans += right - left + 1
	}

	return ans
}
