package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Claim struct {
	ID       string
	Position *Point
	Size     *Size
}

type Point struct {
	X, Y int
}

type Size struct {
	Width, Height int
}

type Fabric [1000][1000]string

func main() {
	input := readInput()
	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}

func part1(input []string) int {
	var fabric Fabric
	for _, line := range input {
		claim := Parse(line)
		mark(&fabric, claim)
	}
	return count(&fabric)
}

func part2(input []string) string {
	var fabric Fabric
	var claims []*Claim
	for _, line := range input {
		c := Parse(line)
		claims = append(claims, c)
		mark(&fabric, c)
	}
	var ans string
	for _, claim := range claims {
		if !isOverWrapped(&fabric, claim) {
			ans = claim.ID
		}
	}
	return ans
}

func isOverWrapped(fabric *Fabric, claim *Claim) bool {
	for i := claim.Position.X; i < claim.Position.X+claim.Size.Width; i++ {
		for j := claim.Position.Y; j < claim.Position.Y+claim.Size.Height; j++ {
			if fabric[i][j] == "X" {
				return true
			}
		}
	}
	return false
}

func mark(fabric *Fabric, claim *Claim) {
	for i := claim.Position.X; i < claim.Position.X+claim.Size.Width; i++ {
		for j := claim.Position.Y; j < claim.Position.Y+claim.Size.Height; j++ {
			if fabric[i][j] == "" {
				fabric[i][j] = claim.ID
			} else {
				fabric[i][j] = "X"
			}
		}
	}
}

func count(fabric *Fabric) int {
	counter := 0
	for _, column := range fabric {
		for _, cell := range column {
			if cell == "X" {
				counter++
			}
		}
	}
	return counter
}

func Parse(line string) *Claim {
	r := regexp.MustCompile(`#(\d.*)\s@\s(\d.*),(\d.*):\s(\d.*)x(\d.*)$`)
	result := r.FindStringSubmatch(line)
	var res []int
	for i := 2; i < len(result); i++ {
		v, _ := strconv.Atoi(result[i])
		res = append(res, v)
	}
	claim := &Claim{
		result[1],
		&Point{res[0], res[1]},
		&Size{res[2], res[3]},
	}
	return claim
}

func readInput() []string {
	file, err := os.Open("./day3input.txt")
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
