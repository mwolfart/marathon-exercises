package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isDigit(b rune) bool {
	return unicode.IsDigit(b)
}

func testNewCharacter(b rune, currentNumStr string) (string, bool) {
	if isDigit(b) {
		return currentNumStr + string(b), true
	} else {
		return currentNumStr, false
	}
}

func getLeadingNumberInLine(line string, start int) string {
	currentNumStr := ""
	for i := start; i >= 0; i-- {
		if !isDigit(rune(line[i])) {
			currentNumStr = line[i+1 : start+1]
			break
		} else if i == 0 {
			currentNumStr = line[:start+1]
		}
	}
	return currentNumStr
}

func getTrailingNumberInLine(line string, start int) string {
	currentNumStr := ""
	for i := start; i < len(line); i++ {
		if !isDigit(rune(line[i])) {
			currentNumStr = line[start:i]
			break
		} else if i == len(line)-1 {
			currentNumStr = line[start:]
		}
	}
	return currentNumStr
}

func getAdjNumbersInLine(line string, initialStart int, initialEnd int) []int {
	start := initialStart
	end := initialEnd
	var numbers []int
	currentNumStr := ""
	var ok bool

	if isDigit(rune(line[start])) && start > 0 {
		currentNumStr = getLeadingNumberInLine(line, start)
		start += 1
	}

	for _, b := range line[start:end] {
		currentNumStr, ok = testNewCharacter(b, currentNumStr)
		if !ok && len(currentNumStr) > 0 {
			n, _ := strconv.Atoi(currentNumStr)
			numbers = append(numbers, n)
			currentNumStr = ""
		}
	}

	if len(currentNumStr) > 0 {
		for i := end; i < len(line); i++ {
			b := rune(line[i])
			currentNumStr, ok = testNewCharacter(b, currentNumStr)
			if !ok || (i+1 == len(line)) {
				n, _ := strconv.Atoi(currentNumStr)
				numbers = append(numbers, n)
				break
			}
		}
	}

	return numbers
}

func getAdjacentNumbers(index int, line string, prevLine string, nextLine string) []int {
	start := int(math.Max(0.0, float64(index-1)))
	end := int(math.Min(float64(len(line)-1), float64(index+2)))
	var numbers []int

	if len(prevLine) > 0 {
		numbers = getAdjNumbersInLine(prevLine, start, end)
	}

	if len(nextLine) > 0 {
		numbers = append(numbers, getAdjNumbersInLine(nextLine, start, end)...)
	}

	if isDigit(rune(line[start])) {
		leadingNumber := getLeadingNumberInLine(line, start)
		ln, _ := strconv.Atoi(leadingNumber)
		numbers = append(numbers, ln)
	}

	if isDigit(rune(line[end-1])) {
		trailingNumber := getTrailingNumberInLine(line, end-1)
		tn, _ := strconv.Atoi(trailingNumber)
		numbers = append(numbers, tn)
	}

	return numbers
}

func getGearRatiosInLine(line string, prevLine string, nextLine string) int {
	remainder := line
	remainderIndex := 0
	lineSum := 0
	var index int

	for len(remainder) > 0 {
		index = strings.Index(remainder, "*")
		if index >= 0 && index < len(remainder)-1 {
			remainder = remainder[index+1:]
		} else {
			remainder = ""
		}

		if index >= 0 {
			lineIndex := index + remainderIndex
			adjacentNumbers := getAdjacentNumbers(lineIndex, line, prevLine, nextLine)

			if len(adjacentNumbers) == 2 {
				lineSum += adjacentNumbers[0] * adjacentNumbers[1]
			}

			remainderIndex += index + 1
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
		sum += getGearRatiosInLine(line, prevLine, nextLine)
		prevLine = line
		line = nextLine
	}

	// Last line
	sum += getGearRatiosInLine(line, prevLine, "")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum of numbers: %d\n", sum)
}
