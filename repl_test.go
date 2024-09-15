package main

import "testing"

func TestCleanWords(t *testing.T) {
	cases := []struct{
		input 	 string
		expected []string
	} {
		{
			input: "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanWords(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match: '%v' vs '%v'", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanWords(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}