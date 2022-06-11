package string_sum

import (
	"testing"
)

func TestStringSum(t *testing.T) {
	in := " "

	got, err := StringSum(in)
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := "2"

	if got != expected {
		t.Errorf("reverseString(), got: %v, expected: %v", got, expected)
	}
}
