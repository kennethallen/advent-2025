package advent_2025

import (
	"log"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/mrtkp9993/go-overflow"
)

type Day02 struct {
	invalid_pairs_sum uint64
	invalid_all_sum   uint64
	seen              mapset.Set[uint64]
}

func (sol *Day02) Process(input string) {
	sol.seen = mapset.NewThreadUnsafeSet[uint64]()
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), ",") {
		extrema := strings.Split(line, "-")
		start, err := strconv.ParseUint(extrema[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.ParseUint(extrema[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		pairs_sum := sol.process_reps(start, end, 2)
		sol.invalid_pairs_sum += pairs_sum
		sol.invalid_all_sum += pairs_sum
		// uint64 up to 20 digits long
		for _, reps := range []uint64{3, 5, 7, 11, 13, 17, 19} {
			sol.invalid_all_sum += sol.process_reps(start, end, reps)
		}
	}
}

func (sol *Day02) process_reps(start, end, reps uint64) uint64 {
	/*
		For each power of ten n, invalid codes are (n + 1)x where x is in [n/10, n-1]
		[1     , 99     ]: 11   n, n in [1  , 9  ]
		[100   , 9,999  ]: 101  n, n in [10 , 99 ]
		[10,000, 999,999]: 1,001n, n in [100, 999]

		For triples
		[1        , 999        ]: 111      n, n in [1  , 9  ]
		[1,000    , 999,999    ]: 10,101   n, n in [10 , 99 ]
		[1,000,000, 999,999,999]: 1,001,001n, n in [100, 999]

		Should handle upper bounds more intelligently than just checking for integer overflow.
		Should sum invalid values in blocks at least instead of one by one ((n^2 + n)/2 etc.).
		Should look for better solution to duplicates than `seen` set.
	*/
	var accum, ten_power uint64 = 0, 10
	one_past_max_invalid := pow(ten_power, reps)
	if one_past_max_invalid == 0 {
		return 0
	}
	for start >= one_past_max_invalid {
		ten_power *= 10
		one_past_max_invalid = pow(ten_power, reps)
		if one_past_max_invalid == 0 {
			return 0
		}
	}

	mult := calc_mult(ten_power, reps)
	if mult == 0 {
		return accum
	}
	x := max(((start-1)/mult)+1, ten_power/10)
	invalid := mult * x
	for invalid <= end {
		if sol.seen.Add(invalid) {
			accum += invalid
		}

		x++
		if x == ten_power {
			ten_power *= 10
			mult = calc_mult(ten_power, reps)
			if mult == 0 {
				return accum
			}
		}
		invalid = mult * x
	}
	return accum
}

func pow(base, exp uint64) uint64 {
	var res uint64 = 1
	for range exp {
		var err bool
		res, err = overflow.MulUint64(res, base)
		if err {
			return 0
		}
	}
	return res
}

func calc_mult(ten_power, reps uint64) uint64 {
	var mult, ten_power_power uint64 = 0, 1
	for range reps {
		var err bool
		mult, err = overflow.AddUint64(mult, ten_power_power)
		if err {
			return 0
		}

		ten_power_power, err = overflow.MulUint64(ten_power_power, ten_power)
		if err {
			return 0
		}
	}
	return mult
}

func (sol *Day02) Part1() uint64 { return sol.invalid_pairs_sum }
func (sol *Day02) Part2() uint64 { return sol.invalid_all_sum }
