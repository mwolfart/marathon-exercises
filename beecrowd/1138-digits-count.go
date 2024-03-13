package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func div_floor(a int, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func num_of_digits(value int) int {
	count := 0
	for i := value; i > 0; i = div_floor(i, 10) {
		count++
	}
	return count
}

func get_digits_count_from_one(value int) [10]int {
	var counts [10]int
	limit := pow(10, num_of_digits(value))

	for digit, _ := range counts {
		amount := 0
		for p := 10; p <= limit; p = p * 10 {
			if digit == 0 && p >= 100 {
				amount += (div_floor(value, p) - 1) * (p / 10)
			} else {
				amount += div_floor(value, p) * (p / 10)
			}
			remainder := value % p
			if digit == 0 && p == 10 {
				continue
			} else if remainder >= (digit+1)*(p/10) {
				amount += p / 10
			} else if remainder >= digit*p/10 {
				amount += remainder - (digit * p / 10) + 1
			}
		}
		counts[digit] = amount
	}

	return counts
}

func get_digits_count(min int, max int) [10]int {
	lower := get_digits_count_from_one(min - 1)
	upper := get_digits_count_from_one(max)
	var result [10]int
	for i, _ := range result {
		result[i] = upper[i] - lower[i]
	}
	return result
}

func main() {
	var input = bufio.NewScanner(os.Stdin)
	min, max := 0, 0

	for input.Scan() {
		line := input.Text()
		fmt.Sscanf(line, "%d %d", &min, &max)

		if min == 0 || max == 0 {
			break
		}

		counts := get_digits_count(min, max)
		for i, value := range counts {
			if i == 9 {
				fmt.Printf("%d\n", value)
			} else {
				fmt.Printf("%d ", value)
			}
		}
	}
}
