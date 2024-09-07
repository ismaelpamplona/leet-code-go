package p917_reverse_only_letters

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "ab-cd"
	expected := "dc-ba"
	result1 := reverseOnlyLetters1(input)
	result2 := reverseOnlyLetters2(input)
	if !reflect.DeepEqual(expected, result1) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
	if !reflect.DeepEqual(expected, result2) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result2, expected)
	}
}

func TestCase2(t *testing.T) {
	input := "a-bC-dEf-ghIj"
	expected := "j-Ih-gfE-dCba"
	result1 := reverseOnlyLetters1(input)
	result2 := reverseOnlyLetters2(input)
	if !reflect.DeepEqual(expected, result1) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
	if !reflect.DeepEqual(expected, result2) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result2, expected)
	}
}

func TestCase3(t *testing.T) {
	input := "Test1ng-Leet=code-Q!"
	expected := "Qedo1ct-eeLg=ntse-T!"
	result1 := reverseOnlyLetters1(input)
	result2 := reverseOnlyLetters2(input)
	if !reflect.DeepEqual(expected, result1) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result1, expected)
	}
	if !reflect.DeepEqual(expected, result2) {
		t.Errorf("Test failed! 😞 Got %v, expected %v", result2, expected)
	}
}
