package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
)

func main() {
	p1Result := P1(utils.ParseInput(os.Args[2])[0])
	p2Result := P2(utils.ParseInput(os.Args[2])[0])

	fmt.Println(p1Result)
	fmt.Println(p2Result)
}

func P1(input string) int {
	return uniqueOccurencesIndex(input, 4)
}

func P2(input string) int {
	return uniqueOccurencesIndex(input, 14)
}

func uniqueOccurencesIndex(input string, numUnique int) int {
	tailIndex := numUnique - 1
	// keep track of the occurrences of chars in the window
	inWindow := map[rune]int{}

	for i, char := range input {
		// increment occurrences of head
		inWindow[char] = inWindow[char] + 1

		// only start counting once the tail is in bounds
		if i < tailIndex {
			continue
		}

		// check occurrences of each character in the window
		allCountsSingular := true
		for _, c := range input[i-tailIndex : i] {
			count := inWindow[c]
			if count > 1 {
				allCountsSingular = false
				break
			}
		}

		// return position if all counts == 1
		if allCountsSingular {
			return i + 1
		}

		// decrement occurrences of tail
		charCount := inWindow[rune(input[i-tailIndex])]
		inWindow[rune(input[i-tailIndex])] = charCount - 1
	}

	return 0
}
