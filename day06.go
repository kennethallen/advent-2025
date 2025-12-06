package advent_2025

import (
	"log"
	"strconv"
	"strings"
)

type Day06 struct {
	grand_total uint64
}

func (sol *Day06) Process(input string) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	ops := lines[len(lines)-1]
	operands := make([]uint64, 0, len(lines)-1)
	for i := 0; i < len(ops); {
		end := i + 1
		for ; end < len(ops); end++ {
			if ops[end] != ' ' {
				end--
				break
			}
		}
		for _, line := range lines[:len(lines)-1] {
			operand, err := strconv.ParseUint(strings.Trim(line[i:end], " "), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			operands = append(operands, operand)
		}
		sol.grand_total += eval(ops[i], operands)

		i = end + 1
		operands = operands[:0]
	}
}

func eval(op byte, operands []uint64) uint64 {
	switch op {
	case '+':
		var sum uint64
		for _, operand := range operands {
			sum += operand
		}
		return sum
	case '*':
		var prod uint64 = 1
		for _, operand := range operands {
			prod *= operand
		}
		return prod
	default:
		panic("Unknown op")
	}
}

func (sol *Day06) Part1() uint64 { return sol.grand_total }
func (sol *Day06) Part2() uint64 { return 0 }
