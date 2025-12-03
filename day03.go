package advent_2025

import (
	"strings"
)

type Day03 struct {
	p1, p2 uint64
}

func (sol *Day03) Process(input string) {
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n") {
		// Init turning on first 12 batteries
		maxes12 := [12]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		// For each jth battery, starting at 1 and while we have 12 digits after us...
		j := 1
		for ; j < len(line)-len(maxes12)+1; j++ {
			// For each position n in the final number...
			for n, i := range maxes12 {
				// End checks if this digit is already in use
				if i == j {
					break
				}
				// Consider substituting the jth battery for the ith battery
				if line[j] > line[i] { // Comparing ASCII codepoints
					// Sub in our new digit and reinit subsequent digit positions to use the next positions
					for offset := 0; n+offset < len(maxes12); offset++ {
						maxes12[n+offset] = j + offset
					}
					break
				}
			}
		}
		// We know at this point the top 2 of the 12 digits are what we would have picked for a 2-battery solution
		maxes2 := [2]int{maxes12[0], maxes12[1]}
		// As above except we are running out of digits...
		for ; j < len(line); j++ {
			// So only consider replacing later digits
			for n := j - len(line) + len(maxes12); n < len(maxes12); n++ {
				i := maxes12[n]
				if i == j {
					break
				}
				if line[j] > line[i] {
					for offset := 0; n+offset < len(maxes12); offset++ {
						maxes12[n+offset] = j + offset
					}
					break
				}
			}
			// Repeat for the 2-battery solution
			for n := max(0, j-len(line)+len(maxes2)); n < len(maxes2); n++ {
				i := maxes2[n]
				if i == j {
					break
				}
				if line[j] > line[i] {
					for offset := 0; n+offset < len(maxes2); offset++ {
						maxes2[n+offset] = j + offset
					}
					break
				}
			}
		}

		sol.p1 += joltage(line, maxes2[:])
		sol.p2 += joltage(line, maxes12[:])
	}
}

func joltage(line string, digit_idxs []int) uint64 {
	var sum uint64
	for _, digit_idx := range digit_idxs {
		sum = sum*10 + uint64(line[digit_idx]-'0')
	}
	return sum
}

func (sol *Day03) Part1() uint64 { return sol.p1 }
func (sol *Day03) Part2() uint64 { return sol.p2 }
