package p1413_minimum_value_to_get_positive_step_by_step_sum

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{-3, 2, -3, 4, 2}
	expected := 5
	result1 := minStartValue(input)
	if result1 != expected {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}

func TestCase2(t *testing.T) {
	input := []int{1, 2}
	expected := 1
	result1 := minStartValue(input)
	if result1 != expected {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}

func TestCase3(t *testing.T) {
	input := []int{1, -2, -3}
	expected := 5
	result1 := minStartValue(input)
	if result1 != expected {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
}
