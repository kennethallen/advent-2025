package advent_2025

import (
	"errors"
	"log"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Day10 struct {
	min_presses uint64
}

type MachineParse struct {
	Lights   []Light  `parser:"'[' (@'.' | @'#')* ']'"`
	Buttons  []Button `parser:"('(' @@* ')')*"`
	Joltages []int    `parser:"'{' @Int (',' @Int)* '}'"`
}

type Light bool

func (light *Light) Capture(values []string) error {
	switch values[0] {
	default:
		return errors.New("unexpected light value")
	case "#":
		*light = true
	case ".":
	}
	return nil
}

type Button struct {
	LightsToggled []int `parser:"@Int (',' @Int)*"`
}

func (sol *Day10) Process(input string) {
	parser, err := participle.Build[MachineParse]()
	if err != nil {
		log.Fatal(err)
	}
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n") {
		ast, err := parser.ParseString("", line)
		if err != nil {
			log.Fatal(err)
		}

		var target uint64
		for i, on := range ast.Lights {
			if on {
				target ^= 1 << i
			}
		}

		buttons := make([]uint64, 0, len(ast.Buttons))
		for _, button_ast := range ast.Buttons {
			var button uint64
			for _, i := range button_ast.LightsToggled {
				button ^= 1 << i
			}
			buttons = append(buttons, button)
		}

		var min_presses uint64
		for ; !recurse(target, 0, min_presses, buttons); min_presses++ {
		}
		sol.min_presses += min_presses
	}
}

func recurse(target, state, remaining_presses uint64, buttons []uint64) bool {
	if target == state && remaining_presses == 0 {
		return true
	}
	if remaining_presses > uint64(len(buttons)) || len(buttons) == 0 {
		return false
	}
	return recurse(target, state^buttons[0], remaining_presses-1, buttons[1:]) ||
		recurse(target, state, remaining_presses, buttons[1:])
}

func (sol *Day10) Part1() uint64 { return sol.min_presses }
func (sol *Day10) Part2() uint64 { return 0 }
