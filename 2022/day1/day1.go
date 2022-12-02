package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines) // scan line by line

	currentCals := 0
	maxCals := 0

	// elfCals := make(map[int]int)
	// elfIndex := 0

	// elfCals[elfIndex] = 0

	elfCals := make([]int, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text() // the line
		if line == "" {
			if currentCals > maxCals {
				maxCals = currentCals
			}
			elfCals = append(elfCals, currentCals)
			currentCals = 0
		} else {
			intLine, err := strconv.Atoi(line)
			check(err)
			currentCals += intLine
			// elfCals[elfIndex] += intLine
		}
	}
	fmt.Println(maxCals)
	fmt.Println(elfCals)

	sort.Slice(elfCals, func(i, j int) bool {
		return elfCals[i] > elfCals[j]
	})

	topThreeSum := 0
	for _, v := range elfCals[0:3] {
		topThreeSum += v
	}
	fmt.Println(topThreeSum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
