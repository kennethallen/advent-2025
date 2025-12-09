package advent_2025

import (
	"testing"
)

func TestDay09(t *testing.T) {
	RunTests[Day09](t, []TestCase[uint64]{
		{input_file: "09.ex.txt", solutions: [2]uint64{50, 0}},
		{input_file: "09.txt", solutions: [2]uint64{4771508457, 0}},
	})
}
