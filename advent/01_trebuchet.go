package main

import (
	"bufio"
	"fmt"
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

func replaceDigits(line string) string {
	r := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	s := r.Replace(line)
	return s
}

func extractNumber(line string) int {
	i := 0
	var first, last byte
	for first = line[i]; i < len(line) && !unicode.IsDigit(rune(first)); i++ {
		first = line[i]
	}
	if !unicode.IsDigit(rune(first)) {
		return 0
	}
	j := len(line) - 1
	for last = line[j]; j >= i-1 && !unicode.IsDigit(rune(last)); j-- {
		last = line[j]
	}
	number, err := strconv.Atoi(string(first) + string(last))
	check(err)
	return number
}

func sum_all(slc []int) int {
	sum := 0
	for _, value := range slc {
		sum += value
	}
	return sum
}

func main() {
	file, err := os.Open("01.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		updated := replaceDigits(line)
		number := extractNumber(updated)
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Sum of numbers: %d\n", sum_all(numbers))
}
