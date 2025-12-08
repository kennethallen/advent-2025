package advent_2025

import (
	"testing"
)

func TestDay08(t *testing.T) {
	RunTests[Day08](t, []TestCase[uint64]{
		{input_file: "08.ex.txt", solutions: [2]uint64{40, 0}},
		{input_file: "08.txt", solutions: [2]uint64{29406, 0}},
	})
}
