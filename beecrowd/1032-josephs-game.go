package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func is_prime(n int) bool {
	sqrt := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func get_survivor(n int) int {
	arr := make([]int, n)
	for i, _ := range arr {
		arr[i] = i + 1
	}
	prime := 2
	counter := -1
	for len(arr) > 1 {
		counter = (counter + prime) % len(arr)

		arr = append(arr[:counter], arr[counter+1:]...)

		if counter == 0 {
			counter = len(arr) - 1
		} else {
			counter--
		}

		if prime == 2 {
			prime++
		} else {
			prime += 2
			for !is_prime(prime) {
				prime += 2
			}
		}
	}

	return arr[0]
}

func main() {
	var input = bufio.NewScanner(os.Stdin)
	n_people := 0

	for input.Scan() {
		line := input.Text()
		fmt.Sscanf(line, "%d", &n_people)

		if n_people == 0 {
			break
		}

		fmt.Println(get_survivor(n_people))
	}
}
