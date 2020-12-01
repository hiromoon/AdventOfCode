package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("part1: ", part1())
	fmt.Println("part2: ", part2())
}

func part2() int {
	ans := []int{}
	for c := byte('a'); c <= byte('z'); c++ {
		filePath := fmt.Sprintf("../input/%s.txt", "day5")
		file, err := os.Open(filePath)
		if err != nil {
			panic(err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		buf := []byte{}
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(err.Error())
			}

			tmp := []byte{}
			for _, b := range scanner.Bytes() {
				if c != b && bytes.ToUpper(([]byte{c}))[0] != b {
					tmp = append(tmp, b)
				}
			}
			buf = compress(tmp, buf)
		}
		ans = append(ans, len(buf))
		scanner = nil
	}
	return findMin(ans)
}

func findMin(ans []int) int {
	min := 0
	for _, l := range ans {
		if min == 0 || min > l {
			min = l
		}
	}

	return min
}

func part1() int {
	filePath := fmt.Sprintf("../input/%s.txt", "day5")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := []byte{}
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
		buf = compress(scanner.Bytes(), buf)
	}
	return len(buf)
}

func compress(readedBytes []byte, buf []byte) []byte {
	for _, b := range readedBytes {
		if len(buf) == 0 {
			buf = append(buf, b)
			continue
		}
		lastIndex := len(buf) - 1
		last := []byte{buf[lastIndex]}
		if last[0] == b {
			buf = append(buf, b)
		} else if (last[0] >= byte('a') && (bytes.ToUpper(last))[0] == b) ||
			(last[0] <= byte('Z') && (bytes.ToLower(last))[0] == b) {
			buf = buf[:lastIndex]
		} else {
			buf = append(buf, b)
		}
	}
	return buf
}
