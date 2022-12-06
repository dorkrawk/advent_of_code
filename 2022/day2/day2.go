package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines) // scan line by line

	their_rps_value := map[string]int{"A": 1, "B": 2, "C": 3} // rock, paper, scissors
	my_rps_value := map[string]int{"X": 1, "Y": 2, "Z": 3}    // rock, paper, scissors

	total_score := 0

	for fileScanner.Scan() {
		line := fileScanner.Text() // the line
		throws := strings.Split(line, " ")
		outcome_score := rps_outcome(their_rps_value[throws[0]], my_rps_value[throws[1]])
		// fmt.Println("******")
		// fmt.Println(outcome_score)
		// fmt.Println(my_rps_value[throws[1]])
		// fmt.Println("------")
		total_score += (outcome_score + my_rps_value[throws[1]])
	}
	fmt.Println(total_score)
}

func rps_outcome(them, me int) int {
	switch pmod(them-me, 3) {
	case 0:
		return 3 // draw
	case 2:
		return 6 // win
	case 1:
		return 0 // loss
	}
	return 0
}

func pmod(v, m int) int {
	mod := v % m
	if v < 0 {
		return mod + m
	} else {
		return mod
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
