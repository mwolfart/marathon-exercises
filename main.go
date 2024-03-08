package main

import "fmt"

func sum2(s []int) int {
	sum := 0
	for _, item := range s {
		sum += item
	}
	return sum
}

func sum1(s []int, c chan int) {
	sum := 0
	for _, item := range s {
		sum += item
	}
	c <- sum
}

func main() {
	s := []int{2, 3, 4, 5, 5, 6, 7, 8}
	ch := make(chan int)

	go sum1(s[:len(s)/2], ch)
	go sum1(s[len(s)/2:], ch)

	x, y, z := <-ch, <-ch, <-ch

	fmt.Println(x, y, z, x+y)
}
