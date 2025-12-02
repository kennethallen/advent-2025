package advent_2025

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase[O any] struct {
	input_file string
	solutions  [2]O
}

type Solver[O any] interface {
	Process(string)
	Part1() O
	Part2() O
}

func RunTests[S any, O comparable, SP interface {
	*S
	Solver[O]
}](t *testing.T, cases []TestCase[O]) {
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
			var skip O
			if sol != skip {
				var res O
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
