package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// LineLog is...
type LineLog struct {
	Time    *time.Time
	Message string
}

// Logs is ...
type Logs map[int][60]int

func main() {
	input := readInput("day4")
	var lineLogs []*LineLog
	for _, line := range input {
		lineLogs = append(lineLogs, parse(line))
	}
	sort.Slice(lineLogs, func(i, j int) bool { return lineLogs[i].Time.Sub(*lineLogs[j].Time) < 0 })
	groupedLogs := group(lineLogs)
	fmt.Println("part1: ", part1(groupedLogs))
	fmt.Println("part2: ", part2(groupedLogs))
}

func part1(logs Logs) int {
	id := findMaxID(logs)
	time := findMaxTime(logs[id])
	return id * time
}

func part2(logs Logs) int {
	var max int
	var maxID int
	var maxTime int
	for id, values := range logs {
		time := findMaxTime(values)
		if values[time] > max {
			max = values[time]
			maxTime = time
			maxID = id
		}
	}

	return maxTime * maxID
}

func findMaxTime(times [60]int) int {
	time := 0
	max := 0
	for t, count := range times {
		if max < count {
			max = count
			time = t
		}
	}
	return time
}

func findMaxID(logs Logs) int {
	var id int
	max := 0
	for i, log := range logs {
		sum := 0
		for _, l := range log {
			sum += l
		}

		if sum > max {
			id = i
			max = sum
		}
	}
	return id
}

func group(lineLogs []*LineLog) Logs {
	var logs Logs
	logs = map[int][60]int{}

	var logID int
	var asleepTime *time.Time
	r := regexp.MustCompile("Guard #(.*) begins shift")
	for _, line := range lineLogs {
		switch line.Message {
		case "falls asleep":
			asleepTime = line.Time
		case "wakes up":
			log := logs[logID]
			for i := asleepTime.Minute(); i < line.Time.Minute(); i++ {
				log[i]++
			}
			logs[logID] = log
		default:
			result := r.FindStringSubmatch(line.Message)
			id, err := strconv.Atoi(result[1])
			if err != nil {
				panic(err.Error())
			}

			logID = id
		}
	}
	return logs
}

func parse(line string) *LineLog {
	r := regexp.MustCompile(`\[(.*)\] (.*)`)
	result := r.FindStringSubmatch(line)
	t, e := time.Parse("2006-01-02 15:04", result[1])
	if e != nil {
		panic(e.Error())
	}
	return &LineLog{
		&t,
		result[2],
	}
}
