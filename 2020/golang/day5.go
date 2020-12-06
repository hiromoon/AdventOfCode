package main

import (
	"fmt"
	"sort"
	"strings"
)

func parseLine(input string) []string {
	return strings.Split(input, "")
}

func generateRange(min, max int) []int {
	result := []int{}
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	return result
}

func getSeatID(line []string, rows, columns []int) int {
	if len(line) == 0 {
		return rows[0]*8 + columns[0]
	}
	switch line[0] {
	case "F":
		return getSeatID(line[1:], rows[0:len(rows)/2], columns)
	case "B":
		return getSeatID(line[1:], rows[len(rows)/2:len(rows)], columns)
	case "L":
		return getSeatID(line[1:], rows, columns[0:len(columns)/2])
	case "R":
		return getSeatID(line[1:], rows, columns[len(columns)/2:len(columns)])
	}

	return -1
}

func solve1(input []string) {
	maxID := 0
	for _, i := range input {
		line := parseLine(i)
		seatID := getSeatID(line, generateRange(0, 127), generateRange(0, 7))

		if seatID > maxID {
			maxID = seatID
		}
	}

	fmt.Printf("ans: %d\n", maxID)
}

func find(seats []int) int {
	for i, s := range seats {
		if seats[i+1] == s+1 {
			continue
		}

		return s + 1
	}
	return -1
}

func solve2(input []string) {
	seats := []int{}
	for _, i := range input {
		line := parseLine(i)
		seats = append(seats, getSeatID(line, generateRange(0, 127), generateRange(0, 7)))
	}
	sort.Slice(seats, func(i, j int) bool { return seats[i] < seats[j] })
	fmt.Printf("ans: %d\n", find(seats))
}

func main() {
	input := readInput("day5")

	solve1(input)
	solve2(input)
}
