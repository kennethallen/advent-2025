package advent_2025

import (
	"log"
	"strconv"
	"strings"
)

type Day01 struct {
	zero_stops  int
	zero_passes int
}

func (sol *Day01) Process(input string) {
	lines := strings.Split(input, "\n")
	dial := 50
	for _, line := range lines[:len(lines)-1] {
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch line[:1] {
		case "L":
			if dial > 0 && dial-val <= 0 {
				sol.zero_passes += 1
			}
			dial -= val
		case "R":
			if dial < 0 && dial+val >= 0 {
				sol.zero_passes += 1
			}
			dial += val
		}
		passes := dial / 100
		if passes < 0 {
			passes = -passes
		}
		sol.zero_passes += passes
		dial %= 100
		if dial == 0 {
			sol.zero_stops += 1
		}
	}
}

func (sol *Day01) Part1() int { return sol.zero_stops }
func (sol *Day01) Part2() int { return sol.zero_passes }
