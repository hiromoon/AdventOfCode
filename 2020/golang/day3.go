package main

import (
	"fmt"
	"strings"
)

func counter(input [][]string, right, down int) int {
	j := right
	count := 0

	for i := 0; i < len(input); i = i + down {
		if i == 0 {
			continue
		}
		if input[i][j] == "#" {
			count++
		}

		j = (j + right) % len(input[i])
	}
	return count
}

func solve1(input [][]string) {
	fmt.Printf("ans: %d\n", counter(input, 3, 1))
}

func solve2(input [][]string) {
	fmt.Printf("ans: %d\n", counter(input, 1, 1)*counter(input, 3, 1)*counter(input, 5, 1)*counter(input, 7, 1)*counter(input, 1, 2))
}

func parse(input []string) [][]string {
	var matrix [][]string
	for _, i := range input {
		line := strings.Split(i, "")
		matrix = append(matrix, line)
	}

	return matrix
}

func main() {
	input := parse(readInput("day3"))

	solve1(input)
	solve2(input)
}
