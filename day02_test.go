package advent_2025

import (
	"testing"
)

func TestDay02(t *testing.T) {
	RunTests[Day02](t, []TestCase{
		{input_file: "02.ex.txt", solutions: [2]int{1227775554, 4174379265}},
		{input_file: "02.txt", solutions: [2]int{12599655151, 20942028255}},
	})
}
