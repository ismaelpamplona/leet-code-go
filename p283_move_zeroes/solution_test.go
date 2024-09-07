package p283_move_zeroes

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(input)
	if reflect.DeepEqual(input, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{1, 2, 3, 6}
	expected := []int{0}
	moveZeroes(input)
	if reflect.DeepEqual(input, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input, expected)
	}
}
