package main

import "testing"

type TestCase struct {
	input  string
	wanted string
}

var tests = []TestCase{
	{"1234567890", "1234567890"},
	{"123 456 7891", "1234567891"},
	{"(123) 456 7892", "1234567892"},
	{"(123) 456-7893", "1234567893"},
	{"123-456-7894", "1234567894"},
	{"(123)456-7892", "1234567892"},
}

func TestNormalizeSimple(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalizeSimple(tc.input)
			if actual != tc.wanted {
				t.Errorf("got: %s, required: %s", actual, tc.wanted)
			}
		})
	}
}

func TestNormalizeRegexp(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalizeRegexp(tc.input)
			if actual != tc.wanted {
				t.Errorf("got: %s, required: %s", actual, tc.wanted)
			}
		})
	}
}
