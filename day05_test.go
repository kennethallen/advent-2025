package advent_2025

import (
	"testing"
)

func TestDay05(t *testing.T) {
	RunTests[Day05](t, []TestCase[uint]{
		{input_file: "05.ex.txt", solutions: [2]uint{3, 0}},
		{input_file: "05.txt", solutions: [2]uint{726, 0}},
	})
}
