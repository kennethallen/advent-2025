package advent_2025

import (
	"testing"
)

func TestDay05(t *testing.T) {
	RunTests[Day05](t, []TestCase[uint64]{
		{input_file: "05.ex.txt", solutions: [2]uint64{3, 14}},
		{input_file: "05.txt", solutions: [2]uint64{726, 354226555270043}},
	})
}
