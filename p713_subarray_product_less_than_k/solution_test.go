package p713_subarray_product_less_than_k

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{10, 5, 2, 6}
	k := 100
	expected := 8
	result := numSubarrayProductLessThanK(input, k)
	if result != expected {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{1, 2, 3}
	k := 0
	expected := 0
	result := numSubarrayProductLessThanK(input, k)
	if result != expected {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}
