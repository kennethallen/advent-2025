package advent_2025

import (
	"testing"
)

func TestDay03(t *testing.T) {
	RunTests[Day03](t, []TestCase[uint64]{
		{input_file: "03.ex.txt", solutions: [2]uint64{357, 3121910778619}},
		{input_file: "03.txt", solutions: [2]uint64{17085, 169408143086082}},
	})
}
