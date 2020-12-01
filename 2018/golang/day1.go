package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := readInput("day1")
	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}

func part1(input []string) int {
	var ans int
	for _, v := range input {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		ans += n
	}
	return ans
}

func part2(input []string) int {
	is := convertToIntList(input)

	found := false
	var ans int
	t := []int{0}
	for !found {
		t = calcFrequencyList(t, is)
		found, ans = getFirstDuplicate(t)
	}
	return ans
}

func calcFrequencyList(ans []int, input []int) []int {
	last := ans[len(ans)-1]
	for _, v := range input {
		last += v
		ans = append(ans, last)
	}
	return ans
}

func convertToIntList(input []string) []int {
	var res []int
	for _, v := range input {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, n)
	}
	return res
}

func getFirstDuplicate(input []int) (bool, int) {
	var ans int
	encountered := map[int]bool{}
	found := false
	for _, v := range input {
		if !encountered[v] {
			encountered[v] = true
		} else {
			ans = v
			found = true
			break
		}
	}

	return found, ans
}
