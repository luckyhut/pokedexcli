package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// create test cases
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "oneword   ",
			expected: []string{"oneword"},
		},
		{
			input:    "  hello        two  three   four",
			expected: []string{"hello", "two", "three", "four"},
		},
	}

	// run test cases
	for _, c := range cases {
		// length check
		actual := cleanInput(c.input)
		fmt.Println(len(actual), len(c.expected))
		fmt.Println(actual)
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch")
		}
		// word check
		for i := range actual {
			word := actual[i]
			expectedword := c.expected
			if word != expectedword[i] {
				t.Errorf("word mismatch")
			}
		}
	}
}
