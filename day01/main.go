package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	P1(utils.ParseInput(os.Args[2]))
	P2(utils.ParseInput(os.Args[2]))
}

func P1(input []string) {
	highestCalories := 0
	currentCalories := 0

	for _, val := range input {
		if val == "" {
			if currentCalories > highestCalories {
				highestCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		currentCalories = currentCalories + calories
	}

	fmt.Println(highestCalories)
}

func P2(input []string) {
	var allCalorieCounts []int
	currentCalories := 0

	for _, val := range input {
		if val == "" {
			allCalorieCounts = append(allCalorieCounts, currentCalories)
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		currentCalories = currentCalories + calories
	}

	// account for last line
	if currentCalories != 0 {
		allCalorieCounts = append(allCalorieCounts, currentCalories)
	}

	sort.Ints(allCalorieCounts)

	first := allCalorieCounts[len(allCalorieCounts)-1]
	second := allCalorieCounts[len(allCalorieCounts)-2]
	third := allCalorieCounts[len(allCalorieCounts)-3]
	total := first + second + third

	fmt.Println(total)
}
