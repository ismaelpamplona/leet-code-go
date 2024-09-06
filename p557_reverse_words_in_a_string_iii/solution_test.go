package p557_reverse_words_in_a_string_iii

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input1 := "Let's take LeetCode contest"
	expected1 := "s'teL ekat edoCteeL tsetnoc"
	result := reverseWords(input1)
	if !reflect.DeepEqual(result, expected1) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input1, expected1)
	}
}

func TestCase2(t *testing.T) {
	input1 := "Mr Ding"
	expected1 := "rM gniD"
	result := reverseWords(input1)
	if !reflect.DeepEqual(result, expected1) {
		t.Errorf("Test case 1 failed. Got %v, expected %v", input1, expected1)
	}
}
