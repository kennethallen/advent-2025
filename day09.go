package advent_2025

import (
	"log"
	"strconv"
	"strings"
)

type Day09 struct {
	reds [][2]uint64
}

func (sol *Day09) Process(input string) {
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n") {
		coords := strings.Split(line, ",")
		x, err := strconv.ParseUint(coords[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseUint(coords[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		sol.reds = append(sol.reds, [2]uint64{x, y})
	}
}

func (sol *Day09) Part1() uint64 {
	var area uint64
	for _, a := range sol.reds {
		for _, b := range sol.reds {
			dx := max(a[0], b[0]) - min(a[0], b[0])
			dy := max(a[1], b[1]) - min(a[1], b[1])
			area = max(area, (dx+1)*(dy+1))
		}
	}
	return area
}

func (sol *Day09) Part2() uint64 { return 0 }
