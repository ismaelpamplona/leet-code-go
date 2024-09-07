package p2540_minimum_common_value

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{2, 4}
	expected1 := 2
	result := getCommon(input1, input2)
	if result != expected1 {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input1, expected1)
	}
}

func TestCase2(t *testing.T) {
	input1 := []int{1, 2, 3, 6}
	input2 := []int{2, 3, 4, 5}
	expected1 := 2
	result := getCommon(input1, input2)
	if result != expected1 {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input1, expected1)
	}
}
