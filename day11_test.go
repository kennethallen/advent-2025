package advent_2025

import (
	"testing"
)

func TestDay11(t *testing.T) {
	RunTests[Day11](t, []TestCase[uint64]{
		{input_file: "11.ex1.txt", solutions: [2]uint64{5, 0}},
		{input_file: "11.ex2.txt", solutions: [2]uint64{0, 2}},
		{input_file: "11.txt", solutions: [2]uint64{668, 294310962265680}},
	})
}
