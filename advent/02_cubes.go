package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findMinimums(line string) (int, int, int) {
	idSetSplit := strings.Split(line, ":")
	sets := strings.Split(idSetSplit[1], ";")

	var minReds, minGreens, minBlues int
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, colors := range cubes {
			amountLabelSplit := strings.Split(colors, " ")
			amount, _ := strconv.Atoi(amountLabelSplit[1])
			isRed := amountLabelSplit[2] == "red"
			isGreen := amountLabelSplit[2] == "green"
			isBlue := amountLabelSplit[2] == "blue"

			if isRed && (minReds == 0 || minReds < amount) {
				minReds = amount
			}
			if isGreen && (minGreens == 0 || minGreens < amount) {
				minGreens = amount
			}
			if isBlue && (minBlues == 0 || minBlues < amount) {
				minBlues = amount
			}
		}
	}
	return minReds, minGreens, minBlues
}

func validateGame(line string, maxReds int, maxGreens int, maxBlues int) (int, bool) {
	idSetSplit := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(idSetSplit[0][5:])
	sets := strings.Split(idSetSplit[1], ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, colors := range cubes {
			amountLabelSplit := strings.Split(colors, " ")
			amount, _ := strconv.Atoi(amountLabelSplit[1])
			isRed := amountLabelSplit[2] == "red"
			isGreen := amountLabelSplit[2] == "green"
			isBlue := amountLabelSplit[2] == "blue"
			if (isRed && amount > maxReds) || (isGreen && amount > maxGreens) || (isBlue && amount > maxBlues) {
				return -1, false
			}
		}
	}
	return gameId, true
}

func main() {
	file, err := os.Open("02_cubes.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	sum := 0
	maxReds := 12
	maxGreens := 13
	maxBlues := 14

	for scanner.Scan() {
		line := scanner.Text()
		id, valid := validateGame(line, maxReds, maxGreens, maxBlues)
		// r, g, b := findMinimums(line)
		// sum += r * g * b

		if valid {
			sum += id
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// fmt.Printf("Sum of game IDs: %d\n", sum)
	fmt.Printf("Sum of products: %d\n", sum)
}
