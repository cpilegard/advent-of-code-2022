package main

import (
	"testing"
)

var testStartingStacks = [][]string{
	{"Z", "N"},      // stack 1
	{"M", "C", "D"}, // stack 2
	{"P"},           // stack 3
}

func Test_P1(t *testing.T) {
	testInput := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expected := "CMZ"

	result := P1(testInput, testStartingStacks)
	if result != expected {
		t.Errorf("result: %s does not equal expected: %s", result, expected)
	}
}

func Test_P2(t *testing.T) {
	testInput := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	expected := "MCD"

	result := P2(testInput, testStartingStacks)
	if result != expected {
		t.Errorf("result: %s does not equal expected: %s", result, expected)
	}
}

func Test_stringToCommand(t *testing.T) {
	testInput := "move 1 from 2 to 3"

	expected := command{
		move: 1,
		from: 2,
		to:   3,
	}

	result := stringToCommand(testInput)
	if result != expected {
		t.Errorf("result: %v does not equal expected: %v", result, expected)
	}
}
