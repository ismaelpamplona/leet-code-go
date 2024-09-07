package p2540_minimum_common_value

func getCommon(nums1 []int, nums2 []int) int {

	i, j := 0, 0
	len1, len2 := len(nums1), len(nums2)

	for i < len1 && j < len2 {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return -1
}
