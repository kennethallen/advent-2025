package advent_2025

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct {
	ranges [][2]uint64
	ids    []uint64
}

func (sol *Day05) Process(input string) {
	past_ranges := false
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) == 0 {
			past_ranges = true
		} else if !past_ranges {
			ends := strings.Split(line, "-")
			start, err := strconv.ParseUint(ends[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			end, err := strconv.ParseUint(ends[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			sol.ranges = append(sol.ranges, [2]uint64{start, end})
		} else {
			id, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			sol.ids = append(sol.ids, id)
		}
	}
}

func (sol *Day05) Part1() uint64 {
	var fresh uint64
	for _, id := range sol.ids {
		for _, fresh_range := range sol.ranges {
			if id >= fresh_range[0] && id <= fresh_range[1] {
				fresh++
				break
			}
		}
	}
	return fresh
}
func (sol *Day05) Part2() uint64 {
	slices.SortFunc(sol.ranges, cmp_ranges)
	var cursor, total uint64
	for _, fresh_range := range sol.ranges {
		total -= max(cursor, fresh_range[0])
		cursor = max(cursor, fresh_range[1]+1)
		total += cursor
	}
	return total
}

func cmp_ranges(a, b [2]uint64) int {
	if a[0] < b[0] {
		return -1
	} else if a[0] > b[0] {
		return 1
	} else if a[1] < b[1] {
		return -1
	} else if a[1] > b[1] {
		return 1
	} else {
		return 0
	}
}
