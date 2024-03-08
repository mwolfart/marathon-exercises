package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isSymbol(b rune) bool {
	return (b != '.' && !unicode.IsDigit(b))
}

func isNumberAdjacentToSymbol(curLine string, nextLine string, prevLine string, idxStart int, idxEnd int) bool {
	// TODO
}

func getNumbersInLineWithAdjacentSymbols(curLine string, nextLine string, prevLine string) []int {
	isProcessingNumber := false
	idxStart := -1
	idxEnd := -1

	numbers := make([]int, 0)

	for i, ch := range curLine {
		if unicode.IsDigit(ch) {
			if !isProcessingNumber {
				idxStart = i
				isProcessingNumber = true
			}
		} else {
			if isProcessingNumber {
				idxEnd = i - 1
				isProcessingNumber = false
			}
		}

		if idxStart != -1 && idxEnd != -1 {
			// number found
			isAdjacent := isNumberAdjacentToSymbol(curLine, nextLine, prevLine, idxStart, idxEnd)
			if isAdjacent {
				parsedNumber, _ := strconv.Atoi(curLine[idxStart : idxEnd+1])
				numbers = append(numbers, parsedNumber)
			}

			idxStart = -1
			idxEnd = -1
		}
	}

	return numbers
}

func sum_all(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// armazena primeira e segunda linhas do arquivo
	scanner.Scan()
	prevLine := scanner.Text()
	scanner.Scan()
	line := scanner.Text()

	// process first line
	numbers := getNumbersInLineWithAdjacentSymbols(prevLine, line, "")

	var nextLine string

	// perform loop from 2nd line on
	for scanner.Scan() {
		nextLine = scanner.Text()

		// slice of numbers in that line
		numbersInLine := getNumbersInLineWithAdjacentSymbols(line, nextLine, prevLine)

		// append numbers
		numbers = append(numbers, numbersInLine...)

		// update references
		prevLine = line
		line = nextLine
	}

	// process last line
	numbersInLastLine := getNumbersInLineWithAdjacentSymbols(nextLine, "", line)
	numbers = append(numbers, numbersInLastLine...)

	// numbers stored in `numbers`
	fmt.Println(sum_all(numbers))
}
