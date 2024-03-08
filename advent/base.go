package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("01.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Operations
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// fmt.Printf("Sum of numbers: %d\n", sum_all(numbers))
}
