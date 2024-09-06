package p1004_max_consecutive_ones_iii

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	expected := 6
	result := longestOnes(input, k)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	expected := 10
	result := longestOnes(input, k)
	if expected != result {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result, expected)
	}
}
