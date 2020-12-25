package utils

import (
	"fmt"
	"testing"
)

func TestValidateBrackets(t *testing.T) {
	testCases := []struct {
		name    string
		isValid bool
	}{
		{
			name:    "[()]{}{[()()]()}",
			isValid: true,
		},
		{
			name:    "[(])",
			isValid: false,
		},
		{
			name:    "[({",
			isValid: false,
		},
		{
			name:    ")}]",
			isValid: false,
		},
		{
			name:    "[()]{}{[()()]()}}}}}}}}",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := ValidateBrackets(tc.name)
			if tc.isValid != isValid {
				t.Errorf("expected: %t\actual: %t", tc.isValid, isValid)
			}
		})
	}
}

func TestFix(t *testing.T) {
	testCases := []struct {
		name   string
		expect string
	}{
		{
			name:   "[()]{}{[()()]()}",
			expect: "[()]{}{[()()]()}",
		},
		{
			name:   "[(])",
			expect: "[]",
		},
		{
			name:   "[({",
			expect: "[({})]",
		},
		{
			name:   ")}]",
			expect: "",
		},
		{
			name:   "[()]{}{[()()]()}}}}}}}}",
			expect: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Fix(tc.name)
			fmt.Println(actual)
			if tc.expect != actual {
				t.Errorf("expect: %s\nactual: %s", tc.expect, actual)
			}
		})
	}
}
