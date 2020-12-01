package main

import (
	"bufio"
	"fmt"
	"os"
)

// readInput("day1")
func readInput(day string) []string {
	filePath := fmt.Sprintf("../input/%s.txt", day)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res []string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
		res = append(res, scanner.Text())
	}
	return res
}
