package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
)

func main() {
	P1(utils.ParseInput(os.Args[2]))
	P2(utils.ParseInput(os.Args[2]))
}

func P1(input []string) {
	priorities := generatePriorities()
	totalPriority := 0

	for _, rucksack := range input {
		sackSize := len(rucksack) / 2
		intersection := getIntersection(rucksack[:sackSize], rucksack[sackSize:])[0]
		totalPriority += priorities[rune(intersection)]
	}

	fmt.Println(totalPriority)
}

func P2(input []string) {
	priorities := generatePriorities()
	totalPriority := 0
	currentGroup := []string{}
	for _, sack := range input {
		currentGroup = append(currentGroup, sack)

		if len(currentGroup) == 3 {
			firstIntersection := getIntersection(currentGroup[0], currentGroup[1])
			finalIntersection := getIntersection(firstIntersection, currentGroup[2])
			totalPriority += priorities[rune(finalIntersection[0])]
			currentGroup = []string{}
		}
	}

	fmt.Println(totalPriority)
}

func getIntersection(s1 string, s2 string) string {
	intersection := ""

	// store first string occurences
	chars := map[rune]bool{}
	for _, char := range s1 {
		chars[char] = true
	}

	// find intersections in second string
	for _, char := range s2 {
		if _, ok := chars[char]; ok {
			intersection = intersection + string(char)
		}
	}

	return intersection
}

func generatePriorities() map[rune]int {
	// generate priorities
	priorities := map[rune]int{}
	priority := 1
	for i := 'a'; i <= 'z'; i++ {
		priorities[i] = priority
		priority++
	}
	priority = 27
	for i := 'A'; i <= 'Z'; i++ {
		priorities[i] = priority
		priority++
	}

	return priorities
}
