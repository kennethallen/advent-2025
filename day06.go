package advent_2025

import (
	"strings"
)

type Day06 struct {
	grand_total_ltr uint64
	grand_total_ttb uint64
}

func (sol *Day06) Process(input string) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	op_lines := lines[:len(lines)-1]
	ops := lines[len(lines)-1]
	ltr_operands := make([]uint64, len(op_lines))
	is_mult := false
	var ttb_prod uint64
	for i := 0; ; i++ {
		if i >= len(ops) || ops[i] != ' ' {
			// Finish op
			if is_mult {
				sol.grand_total_ttb += ttb_prod
				var ltr_prod uint64 = 1
				for _, ltr_operand := range ltr_operands {
					ltr_prod *= ltr_operand
				}
				sol.grand_total_ltr += ltr_prod
			} else {
				for _, ltr_operand := range ltr_operands {
					sol.grand_total_ltr += ltr_operand
				}
			}

			if i >= len(ops) {
				break
			}

			// Reset
			is_mult = ops[i] == '*'
			ttb_prod = 1
			for n := range ltr_operands {
				ltr_operands[n] = 0
			}
		}

		var ttb_operand uint64
		any_digit := false
		for n, op_line := range op_lines {
			if op_line[i] == ' ' {
				continue
			}
			any_digit = true
			digit := uint64(op_line[i] - '0')
			ltr_operands[n] = 10*ltr_operands[n] + digit
			ttb_operand = 10*ttb_operand + digit
		}
		if any_digit {
			if is_mult {
				ttb_prod *= ttb_operand
			} else {
				sol.grand_total_ttb += ttb_operand
			}
		}
	}
}

func (sol *Day06) Part1() uint64 { return sol.grand_total_ltr }
func (sol *Day06) Part2() uint64 { return sol.grand_total_ttb }
