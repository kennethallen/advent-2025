package advent_2025

import (
	"testing"
)

func TestDay06(t *testing.T) {
	RunTests[Day06](t, []TestCase[uint64]{
		{input_file: "06.ex.txt", solutions: [2]uint64{4277556, 3263827}},
		{input_file: "06.txt", solutions: [2]uint64{6343365546996, 11136895955912}},
	})
}
