package p209_minimum_size_subarray_sum

import (
	"testing"
)

func TestCase1(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	expected := 2
	result := minSubArrayLen(target, nums)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	target := 4
	nums := []int{1, 4, 4}
	expected := 1
	result := minSubArrayLen(target, nums)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}

func TestCase3(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expected := 0
	result := minSubArrayLen(target, nums)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}
