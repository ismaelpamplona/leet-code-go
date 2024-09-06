package p1480_running_sum_of_1d_array

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 3, 6, 10}
	result1 := runningSum(input)
	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{1, 1, 1, 1, 1}
	expected := []int{1, 2, 3, 4, 5}
	result1 := runningSum(input)
	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}

func TestCase3(t *testing.T) {
	input := []int{3, 1, 2, 10, 1}
	expected := []int{3, 4, 6, 16, 17}
	result1 := runningSum(input)
	if !reflect.DeepEqual(result1, expected) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}
