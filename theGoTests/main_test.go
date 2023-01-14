package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{
		name:     "valid-case",
		dividend: 234.0,
		divisor:  2.0,
		expected: 117.0,
		isErr:    false,
	},
	{
		name:     "invalid-case",
		dividend: 234.0,
		divisor:  0,
		expected: 0,
		isErr:    true,
	},
	{
		name:     "fraction-case",
		dividend: 5.0,
		divisor:  4.5,
		expected: 1.11111111,
		isErr:    false,
	},
}

func TestDivider(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error, but not got one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f ", tt.expected, got)
		}

	}
}

// func TestDivide(t *testing.T) {

// 	_, err := divide(100.0, 99.0)
// 	if err != nil {
// 		t.Error("Not Expected an error, but got one")
// 	}

// }

// func TestNegativeDivide(t *testing.T) {

// 	_, err := divide(100.0, 0)
// 	if err != nil {
// 		t.Error("Not Expected an error, but got one")
// 	}

// }
