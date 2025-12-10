package advent_2025

import (
	"testing"
)

func TestDay10(t *testing.T) {
	RunTests[Day10](t, []TestCase[uint64]{
		{input_file: "10.ex.txt", solutions: [2]uint64{7, 0}},
		{input_file: "10.txt", solutions: [2]uint64{530, 0}},
	})
}
