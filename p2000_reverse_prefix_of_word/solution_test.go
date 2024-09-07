package p2000_reverse_prefix_of_word

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "abcdefd"
	var ch byte = 'd'
	expected := "dcbaefd"
	result := reversePrefix(input, ch)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input, expected)
	}
}

func TestCase2(t *testing.T) {
	input := "xyxzxe"
	var ch byte = 'z'
	expected := "zxyxxe"
	result := reversePrefix(input, ch)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input, expected)
	}
}

func TestCase3(t *testing.T) {
	input := "abcd"
	var ch byte = 'z'
	expected := "abcd"
	result := reversePrefix(input, ch)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input, expected)
	}
}
