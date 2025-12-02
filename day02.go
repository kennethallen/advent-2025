package advent_2025

import (
	"log"
	"strconv"
	"strings"
)

type Day02 struct {
	invalid_sum int
}

func (sol *Day02) Process(input string) {
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), ",") {
		extrema := strings.Split(line, "-")
		start, err := strconv.Atoi(extrema[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(extrema[1])
		if err != nil {
			log.Fatal(err)
		}

		/*
			For each power of ten n, invalid codes are (n + 1)x where x is in [n/10, n-1]
			[0, 99]          : 11n [1, 9]
			[100, 9,999]     : 101n [10, 99]
			[10,000, 999,999]: 1001n [100, 999]
		*/
		ten_power := 10
		for start >= ten_power*ten_power {
			ten_power *= 10
		}

		mult := ten_power + 1
		x := max(((start-1)/mult)+1, ten_power/10)
		invalid := mult * x
		for invalid <= end {
			sol.invalid_sum += invalid

			x += 1
			if x == ten_power {
				ten_power *= 10
				mult = ten_power + 1
			}
			invalid = mult * x
		}
	}
}

func (sol *Day02) Part1() int { return sol.invalid_sum }
func (sol *Day02) Part2() int { return 0 }
