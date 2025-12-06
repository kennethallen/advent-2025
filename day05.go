package advent_2025

import (
	"log"
	"strconv"
	"strings"
)

type Day05 struct {
	fresh uint
}

func (sol *Day05) Process(input string) {
	lines := strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n")
	ranges := make([][2]uint64, 0)
	for line := range lines {
		if len(line) == 0 {
			break
		}
		ends := strings.Split(line, "-")
		start, err := strconv.ParseUint(ends[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.ParseUint(ends[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, [2]uint64{start, end})
	}

	for line := range lines {
		if len(line) == 0 {
			continue
		}
		id, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		for _, fresh_range := range ranges {
			if id >= fresh_range[0] && id <= fresh_range[1] {
				sol.fresh++
				break
			}
		}
	}
}

func (sol *Day05) Part1() uint { return sol.fresh }
func (sol *Day05) Part2() uint { return 0 }
