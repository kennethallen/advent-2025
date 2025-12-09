package advent_2025

import (
	"log"
	"strconv"
	"strings"
)

type Day09 struct {
	area uint64
}

func (sol *Day09) Process(input string) {
	var reds [][2]uint64
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

		for _, other := range reds {
			dx := max(x, other[0]) - min(x, other[0])
			dy := max(y, other[1]) - min(y, other[1])
			sol.area = max(sol.area, (dx+1)*(dy+1))
		}

		reds = append(reds, [2]uint64{x, y})
	}
}

func (sol *Day09) Part1() uint64 { return sol.area }
func (sol *Day09) Part2() uint64 { return 0 }
