package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// test := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines) // scan line by line
	var sum int32
	sum = 0
	for fileScanner.Scan() {
		line := fileScanner.Text() // the line
		dup_value := getDuplicatePriorityValue(line)
		sum += dup_value
	}
	fmt.Println(sum)
}

func getDuplicatePriorityValue(rucksack string) int32 {
	compartment1 := byteSliceToPriority(rucksack[:len(rucksack)/2])
	compartment2 := byteSliceToPriority(rucksack[len(rucksack)/2:])
	compartment1_contents := make(map[int32]int)
	for _, p := range compartment1 {
		compartment1_contents[p] += 1
	}
	for _, p := range compartment2 {
		if compartment1_contents[p] > 0 {
			return p
		}
	}
	return 0
}

func byteSliceToPriority(compartment_contents string) []int32 {
	priority_slice := make([]int32, len(compartment_contents))

	for i, c := range compartment_contents {
		if c < 97 {
			priority_slice[i] = c - 38 // capital letters
		} else {
			priority_slice[i] = c - 96 // lower case letters
		}
	}
	return priority_slice
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
