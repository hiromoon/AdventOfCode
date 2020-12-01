package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := convert(readInput("day1"))
	// solve1(input)
	solve2(input)
}

func solve1(input []int) {
	for i, a := range input {
		if i+1 >= len(input) {
			continue
		}
		for _, b := range input[i+1:] {
			if a+b == 2020 {
				fmt.Printf("answer: %d and %d: %d\n", a, b, a*b)
				os.Exit(0)
			}
		}
	}
}

func solve2(input []int) {
	for i, a := range input {
		if i+1 >= len(input) {
			continue
		}
		for _, b := range input[i+1:] {
			if i+2 >= len(input) {
				continue
			}
			for _, c := range input[i+2:] {
				if a+b+c == 2020 {
					fmt.Printf("answer: %d, %d and %d: %d\n", a, b, c, a*b*c)
					os.Exit(0)
				}
			}
		}
	}
}

func convert(input []string) []int {
	result := []int{}
	for _, s := range input {
		if s == "" {
			break
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, i)
	}

	return result
}
