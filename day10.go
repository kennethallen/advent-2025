package advent_2025

import (
	"errors"
	"log"
	"slices"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Day10 struct {
	asts []MachineParse
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
		sol.asts = append(sol.asts, *ast)
	}
}

func recurse(target, state uint64, remaining_presses int, buttons []uint64) bool {
	if remaining_presses == 0 {
		return target == state
	}
	if remaining_presses > len(buttons) || len(buttons) == 0 {
		return false
	}
	return recurse(target, state^buttons[0], remaining_presses-1, buttons[1:]) ||
		recurse(target, state, remaining_presses, buttons[1:])
}

func recurse_p2(target, state []int, remaining_presses int, buttons []Button) bool {
	//log.Printf("%v %v %v %v", buttons, target, state, remaining_presses)
	if remaining_presses == 0 {
		return slices.Equal(target, state)
	}
	if len(buttons) == 0 {
		return false
	}
	new_state := slices.Clone(state)
	for i := remaining_presses; i >= 0; i-- {
		for _, j := range buttons[0].LightsToggled {
			new_state[j] = state[j] + i
		}
		if recurse_p2(target, new_state, remaining_presses-i, buttons[1:]) {
			return true
		}
	}
	return false
}

func (sol *Day10) Part1() int {
	total := 0
	for _, ast := range sol.asts {
		if len(ast.Lights) > 64 {
			log.Fatal("too many lights for a u64")
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

		min_presses := 0
		for ; !recurse(target, 0, min_presses, buttons); min_presses++ {
		}
		total += min_presses
	}
	return total
}
func (sol *Day10) Part2() int {
	total := 0
	for _, ast := range sol.asts {
		//log.Println(line)
		min_presses := slices.Max(ast.Joltages)
		slices.SortFunc(ast.Buttons, func(a, b Button) int { return len(b.LightsToggled) - len(a.LightsToggled) })
		for ; !recurse_p2(ast.Joltages, make([]int, len(ast.Joltages)), min_presses, ast.Buttons); min_presses++ {
			//log.Println(min_presses)
		}
		//log.Println(min_presses)
		total += min_presses
	}
	return total
}
