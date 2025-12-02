package advent_2025

import (
	"testing"
)

func TestDay01(t *testing.T) {
	RunTests[Day01](t, []TestCase[int]{
		{input_file: "01.ex.txt", solutions: [2]int{3, 6}},
		{input_file: "01.txt", solutions: [2]int{1141, 6634}},
	})
}
