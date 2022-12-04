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
	// p2Result := P2(utils.ParseInput(os.Args[2]))

	fmt.Println(p1Result)
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
