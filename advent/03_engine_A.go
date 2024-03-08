package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNextNumberPosition(line string) (int, int, string) {
	start := -1
	length := -1
	found := false
	for i, b := range line {
		isDigit := unicode.IsDigit(rune(b))
		if !found && isDigit {
			start = i
			found = true
		} else if found && !isDigit {
			length = i - start
			break
		}
	}

	if found && length == -1 {
		length = len(line) - start
	}

	if found {
		end := start + length
		return start, length, line[end:]
	} else {
		return -1, -1, ""
	}
}

func isSymbol(b rune) bool {
	return (b != '.' && !unicode.IsDigit(b))
}

func isNumberAdjacentToSymbol(index int, length int, line string, prevLine string, nextLine string) bool {
	start := int(math.Max(0.0, float64(index-1)))
	end := int(math.Min(float64(len(line)-1), float64(index+length+1)))

	if len(prevLine) > 0 {
		for _, b := range prevLine[start:end] {
			if isSymbol(b) {
				return true
			}
		}
	}

	if len(nextLine) > 0 {
		for _, b := range nextLine[start:end] {
			if isSymbol(b) {
				return true
			}
		}
	}

	if index > 0 && isSymbol(rune(line[start])) {
		return true
	}

	if index < len(line)-1 && isSymbol(rune(line[end-1])) {
		return true
	}

	return false
}

func sumAdjacentInLine(line string, prevLine string, nextLine string) int {
	remainder := line
	remainderIndex := 0
	lineSum := 0
	var index, length int

	for len(remainder) > 0 {
		index, length, remainder = getNextNumberPosition(remainder)

		if index >= 0 && length >= 1 {
			lineIndex := index + remainderIndex
			isAdjacent := isNumberAdjacentToSymbol(lineIndex, length, line, prevLine, nextLine)

			if isAdjacent {
				parsed, _ := strconv.Atoi(line[lineIndex : lineIndex+length])
				lineSum += parsed
			}

			remainderIndex += index + length
		}
	}

	return lineSum
}

func main() {
	file, err := os.Open("03_engine.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)

	sum := 0
	var prevLine, nextLine string

	scanner.Scan()
	line := scanner.Text()

	for scanner.Scan() {
		nextLine = scanner.Text()
		sum += sumAdjacentInLine(line, prevLine, nextLine)
		prevLine = line
		line = nextLine
	}

	// Last line
	sum += sumAdjacentInLine(line, prevLine, "")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum of numbers: %d\n", sum)
}
