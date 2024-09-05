package p643_maximum_average_subarray_i

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75000
	result := findMaxAverage(input, k)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{5}
	k := 1
	expected := 5.00000
	result := findMaxAverage(input, k)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}
