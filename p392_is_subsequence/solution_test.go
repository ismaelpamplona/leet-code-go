package p392_is_subsequence

import (
	"testing"
)

func TestCase1(t *testing.T) {
	s1 := "abc"
	s2 := "ahbgdc"
	expected := true
	result := isSubsequence(s1, s2)
	if result != expected {
		t.Errorf("Test case 1 failed. Got %v, expected %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	s1 := "axc"
	s2 := "ahbgdc"
	expected := false
	result := isSubsequence(s1, s2)
	if result != expected {
		t.Errorf("Test case 1 failed. Got %v, expected %v", result, expected)
	}
}
