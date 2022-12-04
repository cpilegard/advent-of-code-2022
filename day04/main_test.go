package main

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7", // yes
		"6-6,4-6", // yes
		"2-6,4-8",
		"2-6,4-8",
		"12-88,20-75", // yes
		"9-80,10-80",  // yes
	}

	expected := 4

	result := P1(input)
	if result != expected {
		fmt.Errorf("result: %d did not equal expected: %d", result, expected)
	}
}
