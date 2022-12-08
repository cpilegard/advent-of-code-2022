package main

import "testing"

var input = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func Test_P1(t *testing.T) {
	expected := 95437
	result := P1(input)

	if result != expected {
		t.Errorf("result: %d does not equal expected: %d", result, expected)
	}
}

func Test_P2(t *testing.T) {
	expected := 24933642
	result := P2(input)

	if result != expected {
		t.Errorf("result: %d does not equal expected: %d", result, expected)
	}
}
