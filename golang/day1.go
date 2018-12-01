package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()
	var ans int
	for _, v := range input {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		ans += n
	}
	fmt.Println(ans)
}

func readInput() []string {
	file, err := os.Open("./day1input.txt")
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
