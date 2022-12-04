package main

import (
	"advent-of-code-2022/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p1Result := P1(utils.ParseInput(os.Args[2]))
	p2Result := P2(utils.ParseInput(os.Args[2]))

	fmt.Println(p1Result)
	fmt.Println(p2Result)
}

func P1(input []string) int {
	containingPairCount := 0

	for _, pair := range input {
		ranges := strings.Split(pair, ",")
		range1 := ranges[0]
		range2 := ranges[1]

		elf1Low, _ := strconv.Atoi(strings.Split(range1, "-")[0])
		elf1High, _ := strconv.Atoi(strings.Split(range1, "-")[1])
		elf2Low, _ := strconv.Atoi(strings.Split(range2, "-")[0])
		elf2High, _ := strconv.Atoi(strings.Split(range2, "-")[1])

		if (elf1Low <= elf2Low && elf1High >= elf2High) || (elf2Low <= elf1Low && elf2High >= elf1High) {
			containingPairCount++
		}
	}

	return containingPairCount
}

// Establish the first pair as a boundary and check if the second pair overlaps
func P2(input []string) int {
	totalOverlapCount := 0

	for _, pair := range input {
		ranges := strings.Split(pair, ",")
		range1 := ranges[0]
		range2 := ranges[1]

		boundsLow, _ := strconv.Atoi(strings.Split(range1, "-")[0])
		boundsHigh, _ := strconv.Atoi(strings.Split(range1, "-")[1])
		elfLow, _ := strconv.Atoi(strings.Split(range2, "-")[0])
		elfHigh, _ := strconv.Atoi(strings.Split(range2, "-")[1])

		if (elfLow >= boundsLow && elfLow <= boundsHigh) || // low is in bounds
			(elfHigh <= boundsHigh && elfHigh >= boundsLow) || // high is in bounds
			(elfLow < boundsLow && elfHigh > boundsHigh) { // bounds are encapsulated

			totalOverlapCount++
		}
	}

	return totalOverlapCount
}
