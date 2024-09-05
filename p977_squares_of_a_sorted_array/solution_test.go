package p977_squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{-4, -1, 0, 3, 10}
	expected := []int{0, 1, 9, 16, 100}
	result := sortedSquares(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{-7, -3, 2, 3, 11}
	expected := []int{4, 9, 9, 49, 121}
	result := sortedSquares(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", result, expected)
	}
}
