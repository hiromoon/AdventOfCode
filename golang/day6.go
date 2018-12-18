package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Point is ...
type Point struct {
	X, Y int
}

func main() {
	input := readInput("day6")
	fmt.Println("Part1: ", part1(input))
	fmt.Println("Part2: ", part2(input))
}
func part2(input []string) int {
	points := parse(input)
	maxPoint := getMaxPoint(points)
	area := initArea(maxPoint)
	for x := 0; x < maxPoint.X; x++ {
		for y := 0; y < maxPoint.Y; y++ {
			area[x][y] = sumDistances(Point{x, y}, points)
		}
	}

	counter := 0
	for _, row := range area {
		for _, cell := range row {
			if cell < 10000 {
				counter++
			}
		}
	}
	return counter
}

func sumDistances(point Point, points []Point) int {
	res := 0
	for _, p := range points {
		res += getDistance(point, p)
	}
	return res
}

func part1(input []string) int {
	points := parse(input)
	maxPoint := getMaxPoint(points)
	area := initArea(maxPoint)
	for x := 0; x < maxPoint.X; x++ {
		for y := 0; y < maxPoint.Y; y++ {
			area[x][y] = getShortestDistance(Point{x, y}, points)
		}
	}

	ans := getMaxArea(area, len(points))

	return ans
}

func getInfiniteIDs(area [][]int) []int {
	infiniteIDs := map[int]bool{}
	top := area[0]
	for _, id := range top {
		infiniteIDs[id] = true
	}
	bottom := area[len(area)-1]
	for _, id := range bottom {
		infiniteIDs[id] = true
	}

	for _, row := range area {
		infiniteIDs[row[0]] = true
		infiniteIDs[row[len(row)-1]] = true
	}

	res := []int{}
	for key := range infiniteIDs {
		res = append(res, key)
	}
	return res
}

func getMaxArea(area [][]int, points int) int {
	counter := map[int]int{}
	for _, row := range area {
		for _, cell := range row {
			if cell != -1 {
				counter[cell]++
			}
		}
	}

	for _, id := range getInfiniteIDs(area) {
		delete(counter, id)
	}

	max := 0
	for _, v := range counter {
		if max < v {
			max = v
		}
	}
	return max
}

func initArea(point Point) [][]int {
	area := make([][]int, point.X)
	for i := 0; i < point.X; i++ {
		area[i] = make([]int, point.Y)
	}
	return area
}

func getShortestDistance(point Point, points []Point) int {
	distance := 1000
	isDuplicated := false
	index := 0
	for i, p := range points {
		d := getDistance(point, p)
		if d == distance {
			isDuplicated = true
		} else if d < distance {
			isDuplicated = false
			distance = d
			index = i
		}
	}

	if isDuplicated {
		return -1
	}
	return index
}

func getDistance(p1 Point, p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

func parse(input []string) []Point {
	points := []Point{}
	for _, p := range input {
		t := strings.Split(p, ", ")
		x, _ := strconv.Atoi(t[0])
		y, _ := strconv.Atoi(t[1])
		points = append(points, Point{x, y})
	}
	return points
}

func getMaxPoint(points []Point) Point {
	var x int
	var y int
	for _, p := range points {
		if p.X > x {
			x = p.X
		}
		if p.Y > y {
			y = p.Y
		}
	}

	return Point{x, y}
}
