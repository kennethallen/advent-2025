package advent_2025

import (
	"testing"
)

func TestDay10(t *testing.T) {
	RunTests[Day10](t, []TestCase[int]{
		{input_file: "10.ex.txt", solutions: [2]int{7, 33}},
		{input_file: "10.txt", solutions: [2]int{530, 0}},
	})
}
