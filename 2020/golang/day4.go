package main

// 同じことやったつもりだったが、JS だと解けて Go だと解けなかった。。。
import (
	"fmt"
	"strings"
)

func checkValid(line string) bool {
	words := []string{
		"byr:",
		"iyr:",
		"eyr:",
		"hgt:",
		"hcl:",
		"ecl:",
		"pid:",
	}
	for _, w := range words {
		if !strings.Contains(line, w) {
			return false
		}
	}
	return true
}

func solve1(input []string) {
	line := ""
	count := 0
	for _, l := range input {
		if l != "" {
			line = line + " " + l
			continue
		}

		if checkValid(line) {
			count++
		}

		line = ""
	}

	fmt.Printf("ans: %d", count)
}

func main() {
	input := readInput("day4")

	solve1(input)
}
