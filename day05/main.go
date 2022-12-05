package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var startingStacks = [][]string{
	{"N", "S", "D", "C", "V", "Q", "T"},      // stack 1
	{"M", "F", "V"},                          // stack 2
	{"F", "Q", "W", "D", "P", "N", "H", "M"}, // stack 3
	{"D", "Q", "R", "T", "F"},                // stack 4
	{"R", "F", "M", "N", "Q", "H", "V", "B"}, // stack 5
	{"C", "F", "G", "N", "P", "W", "Q"},      // stack 6
	{"W", "F", "R", "L", "C", "T"},           // stack 7
	{"T", "Z", "N", "S"},                     // stack 8
	{"M", "S", "D", "J", "R", "Q", "H", "N"}, // stack 9
}

type command struct {
	move int
	from int
	to   int
}

func main() {
	// p1Result := P1(utils.ParseInput(os.Args[2]), startingStacks)
	p2Result := P2(utils.ParseInput(os.Args[2]), startingStacks)

	// fmt.Println(p1Result)
	fmt.Println(p2Result)
}

func P1(input []string, stacks [][]string) string {
	result := ""

	// parse input into series of commands
	// for each command, execute number of steps
	for _, commandStr := range input {
		cmd := stringToCommand(commandStr)

		for i := 0; i < cmd.move; i++ {
			// pop `from` and push value to `to`
			val := pop(&stacks[cmd.from-1])
			stacks[cmd.to-1] = append(stacks[cmd.to-1], val)
		}
	}

	// get top object for each stack
	for _, stack := range stacks {
		top := stack[len(stack)-1]
		result += top
	}

	return result
}

func P2(input []string, stacks [][]string) string {
	result := ""

	// parse input into series of commands
	// for each command, execute number of steps
	for _, commandStr := range input {
		cmd := stringToCommand(commandStr)

		fromStack := stacks[cmd.from-1]

		// last N values
		valsToMove := fromStack[len(fromStack)-cmd.move:]

		// move to new stack
		stacks[cmd.to-1] = append(stacks[cmd.to-1], valsToMove...)

		// remove from old stack
		stacks[cmd.from-1] = fromStack[:(len(fromStack) - cmd.move)]
	}

	// get top object for each stack
	for _, stack := range stacks {
		top := stack[len(stack)-1]
		result += top
	}

	return result
}

func stringToCommand(input string) command {
	parsedInput := strings.Split(input, " ")
	moveInt, _ := strconv.Atoi(parsedInput[1])
	fromInt, _ := strconv.Atoi(parsedInput[3])
	toInt, _ := strconv.Atoi(parsedInput[5])

	return command{
		move: moveInt,
		from: fromInt,
		to:   toInt,
	}
}

func pop(stack *[]string) string {
	val := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return val
}
