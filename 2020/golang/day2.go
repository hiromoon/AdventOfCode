package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Line struct {
	min       int
	max       int
	charactor string
	password  string
}

func parse(l string) *Line {
	line := &Line{}
	tmp := strings.Split(l, ": ")

	line.password = tmp[1]
	tmp = strings.Split(tmp[0], " ")
	line.charactor = tmp[1]
	tmp = strings.Split(tmp[0], "-")
	var err error
	line.min, err = strconv.Atoi(tmp[0])
	if err != nil {
		panic(err)
	}
	line.max, err = strconv.Atoi(tmp[1])
	if err != nil {
		panic(err)
	}

	return line
}

func (l *Line) containsCount() int {
	count := 0

	for _, c := range strings.Split(l.password, "") {
		if l.charactor == c {
			count++
		}
	}

	return count
}

func solve1(input []string) {
	var lines []*Line
	for _, i := range input {
		lines = append(lines, parse(i))
	}

	count := 0
	for _, l := range lines {
		c := l.containsCount()
		if l.min <= c && c <= l.max {
			count++
		}
	}

	fmt.Printf("ans: %d\n", count)
}

func solve2(input []string) {
	var lines []*Line
	for _, i := range input {
		lines = append(lines, parse(i))
	}

	count := 0
	for _, l := range lines {
		if string(l.password[l.min-1]) == l.charactor && string(l.password[l.max-1]) == l.charactor {
			continue
		} else if string(l.password[l.min-1]) == l.charactor {
			count++
		} else if string(l.password[l.max-1]) == l.charactor {
			count++
		}
	}
	fmt.Printf("ans: %d\n", count)
}

func main() {
	input := readInput("day2")
	solve1(input)
	solve2(input)
}
