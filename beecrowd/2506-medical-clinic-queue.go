package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Patient struct {
	Hour     int
	Minute   int
	Critical int
}

func get_num_of_criticals(ps []*Patient) int {
	INIT_HOUR := 7
	total_patients_critical := 0
	ellapsed_time := 0

	for _, p := range ps {
		arrival_mins := (p.Hour-INIT_HOUR)*60 + p.Minute
		limit := arrival_mins + p.Critical

		if ellapsed_time < arrival_mins {
			ellapsed_time = int(math.Ceil(float64(arrival_mins)/30)) * 30
		}

		if ellapsed_time > limit {
			total_patients_critical++
		}
		ellapsed_time += 30
	}
	return total_patients_critical
}

func main() {
	var input = bufio.NewScanner(os.Stdin)
	n_patients := 0
	patients := make([]*Patient, 0)

	for input.Scan() {
		line := input.Text()
		if n_patients == 0 {
			fmt.Sscanf(line, "%d", &n_patients)
		} else {
			p := &Patient{}
			fmt.Sscanf(line, "%d %d %d", &p.Hour, &p.Minute, &p.Critical)
			patients = append(patients, p)
			n_patients--
		}
		if n_patients == 0 && len(patients) > 0 {
			fmt.Println(get_num_of_criticals(patients))
			patients = make([]*Patient, n_patients)
		}
	}
}
