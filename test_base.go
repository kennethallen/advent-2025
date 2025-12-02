package advent_2025

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input_file string
	solutions  [2]int
}

type Solver interface {
	Process(string)
	Part1() int
	Part2() int
}

func RunTests[S any, SP interface {
	*S
	Solver
}](t *testing.T, cases []TestCase) {
	for _, test := range cases {
		input, err := os.ReadFile("data/" + test.input_file)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				t.Logf("Skipping missing input file: %s", test.input_file)
				continue
			}
			log.Fatal(err)
		}

		var inner S
		solver := SP(&inner)
		solver.Process(string(input))
		for i, sol := range test.solutions {
			if sol != 0 {
				var res int
				switch i {
				case 0:
					res = solver.Part1()
				case 1:
					res = solver.Part2()
				}
				assert.Equal(t, sol, res, "Input %s, part %d", test.input_file, i+1)
			}
		}
	}
}
