package advent_2025

import (
	"testing"
)

func TestDay04(t *testing.T) {
	RunTests[Day04](t, []TestCase[uint]{
		{input_file: "04.ex.txt", solutions: [2]uint{13, 43}},
		{input_file: "04.txt", solutions: [2]uint{1411, 8557}},
	})
}
