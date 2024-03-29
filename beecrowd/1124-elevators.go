package main

import (
	"fmt"
	"math"
)

func fits(l, c, r1, r2 float64) string {
	if l < 2*r1 || l < 2*r2 || c < 2*r1 || c < 2*r2 {
		return "N"
	}
	cat_x := l - r1 - r2
	cat_y := c - r1 - r2
	if r1+r2 > math.Sqrt(math.Pow(cat_x, 2)+math.Pow(cat_y, 2)) {
		return "N"
	}
	return "S"
}

func main() {
	var l, c, r1, r2 int
	for true {
		fmt.Scanf("%d %d %d %d", &l, &c, &r1, &r2)
		if l == 0 && c == 0 {
			break
		}
		fmt.Printf("%s\n", fits(float64(l), float64(c), float64(r1), float64(r2)))
	}
}
