package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// First Round

// A = opponent rock
// B = opponent paper
// C = opponent scissors

// X = player rock
// Y = player paper
// Z = player scissors

var playerPoints = map[string]int{
	"X": 1, // rock
	"Y": 2, // paper
	"Z": 3, // scissors
}

// [opponent play][player play]result
var roundResult = map[string]map[string]int{
	"A": {
		"X": 3,
		"Y": 6,
		"Z": 0,
	},
	"B": {
		"X": 0,
		"Y": 3,
		"Z": 6,
	},
	"C": {
		"X": 6,
		"Y": 0,
		"Z": 3,
	},
}

// Second Round

var secondRoundResult = map[string]int{
	"X": 0, // need to lose
	"Y": 3, // need to draw
	"Z": 6, // need to win
}

// [opponent play][result]player points
var secondRoundPlayPoints = map[string]map[string]int{
	"A": { // rock
		"X": playerPoints["Z"], // lose, scissors
		"Y": playerPoints["X"], // draw, rock
		"Z": playerPoints["Y"], // win, paper
	},
	"B": { // paper
		"X": playerPoints["X"], // lose, rock
		"Y": playerPoints["Y"], // draw, paper
		"Z": playerPoints["Z"], // win, scissors
	},
	"C": { // scissors
		"X": playerPoints["Y"], // lose, paper
		"Y": playerPoints["Z"], // draw, scissors
		"Z": playerPoints["X"], // win, rock
	},
}

func main() {
	filepath := os.Args[2]
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// d2p1(file)
	d2p2(file)
}

func d2p1(file *os.File) {
	totalPoints := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		moves := strings.Split(text, " ")
		opponentMove := moves[0]
		playerMove := moves[1]

		roundPoints := roundResult[opponentMove][playerMove] + playerPoints[playerMove]

		totalPoints = totalPoints + roundPoints
	}

	fmt.Println(totalPoints)
}

func d2p2(file *os.File) {
	totalPoints := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		moves := strings.Split(text, " ")
		opponentMove := moves[0]
		result := moves[1]

		// roundPoints := roundResult[opponentMove][playerMove] + playerPoints[playerMove]
		roundPoints := secondRoundResult[result] + secondRoundPlayPoints[opponentMove][result]

		totalPoints = totalPoints + roundPoints
	}

	fmt.Println(totalPoints)
}
