package advent_2025

import (
	"testing"
)

func TestDay07(t *testing.T) {
	RunTests[Day07](t, []TestCase[uint]{
		{input_file: "07.ex.txt", solutions: [2]uint{21, 0}},
		{input_file: "07.txt", solutions: [2]uint{1499, 0}},
	})
}
