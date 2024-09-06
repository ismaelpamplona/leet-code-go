package p2090_k_radius_subarray_averages

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	expected := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}
	result1 := getAverages(input, k)
	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{100000}
	k := 0
	expected := []int{100000}
	result1 := getAverages(input, k)
	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}

func TestCase3(t *testing.T) {
	input := []int{8}
	k := 100000
	expected := []int{-1}
	result1 := getAverages(input, k)
	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}
