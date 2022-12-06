package main

import "testing"

func Test_P1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			result := P1(tC.input)

			if result != tC.expected {
				t.Errorf("result: %d does not equal expected: %d", result, tC.expected)
			}
		})
	}
}

func Test_P2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			result := P2(tC.input)

			if result != tC.expected {
				t.Errorf("result: %d does not equal expected: %d", result, tC.expected)
			}
		})
	}
}
