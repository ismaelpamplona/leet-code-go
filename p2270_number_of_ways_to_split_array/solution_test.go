package p2270_number_of_ways_to_split_array

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{10, 4, -8, 7}
	expected := 2
	result1 := waysToSplitArray1(input)
	result2 := waysToSplitArray2(input)
	if expected != result1 {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
	if expected != result2 {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result2, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{2, 3, 1, 0}
	expected := 2
	result1 := waysToSplitArray1(input)
	result2 := waysToSplitArray2(input)
	if expected != result1 {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
	if expected != result2 {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result2, expected)
	}
}
