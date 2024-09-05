package p344_reverse_string

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input1 := []byte{'h', 'e', 'l', 'l', 'o'}
	expected1 := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(input1)
	if !reflect.DeepEqual(input1, expected1) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input1, expected1)
	}
}

func TestCase2(t *testing.T) {
	input2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expected2 := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(input2)
	if !reflect.DeepEqual(input2, expected2) {
		t.Errorf("Test case 2 failed. Got %v, expected %v", input2, expected2)
	}
}
