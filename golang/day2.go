package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := readInput("day2")
	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}

func part1(input []string) int {
	countTwo := 0
	countThree := 0
	for _, i := range input {
		two, three := countAnyLetter(i)
		if two > 0 {
			countTwo++
		}
		if three > 0 {
			countThree++
		}
	}
	return countTwo * countThree
}

func part2(input []string) string {
	d := []string{}
	s := []string{}
	for i, s1 := range input {
		for j := i + 1; j < len(input); j++ {
			a := diff(s1, input[j])
			if len(a) == 1 {
				d = append(d, a)
				s = append(s, s1)
			}
		}
	}
	fmt.Println("diff: ", s)
	ans := []string{}
	for i, a := range s {
		ans = append(ans, strings.Replace(a, d[i], "", 1))
	}
	return strings.Join(ans, "")
}

func diff(src string, dst string) string {
	var tmp []string
	s := strings.Split(src, "")
	d := strings.Split(dst, "")
	for i, v := range s {
		if d[i] != v {
			tmp = append(tmp, v)
		}
	}
	return strings.Join(tmp, "")
}

func sortChar(str string) string {
	ss := strings.Split(str, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func countAnyLetter(str string) (int, int) {
	two := 0
	three := 0
	for _, s := range strings.Split(uniq(str), "") {
		switch strings.Count(str, s) {
		case 2:
			two++
		case 3:
			three++
		default:
			// 何もしない
		}
	}
	return two, three
}

func uniq(str string) string {
	var ss []string
	encountered := map[string]bool{}
	for _, s := range strings.Split(str, "") {
		if !encountered[s] {
			encountered[s] = true
			ss = append(ss, s)
		}
	}

	return strings.Join(ss, "")
}
