package advent_2025

import (
	"strings"
)

type Day04 struct {
	grid [][]bool
}

func (sol *Day04) Process(input string) {
	sol.grid = make([][]bool, 0)
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n") {
		row := make([]bool, 0, len(line))
		for _, char := range line {
			row = append(row, char == '@')
		}
		sol.grid = append(sol.grid, row)
	}
}

func (sol *Day04) Part1() uint {
	var open uint
	for y, row := range sol.grid {
		for x := range row {
			if sol.accessible_paper(y, x) {
				open++
			}
		}
	}
	return open
}

func (sol *Day04) Part2() uint {
	var removed uint
	for {
		acted := false
		for y, row := range sol.grid {
			for x := range row {
				if sol.accessible_paper(y, x) {
					removed++
					acted = true
					sol.grid[y][x] = false
				}
			}
		}
		if !acted {
			return removed
		}
	}
}

func (sol *Day04) accessible_paper(y, x int) bool {
	if !sol.grid[y][x] {
		return false
	}
	adj := 0
	for _, neighbor := range [][2]int{
		{y - 1, x - 1},
		{y - 1, x},
		{y - 1, x + 1},
		{y, x - 1},
		{y, x + 1},
		{y + 1, x - 1},
		{y + 1, x},
		{y + 1, x + 1},
	} {
		ny := neighbor[0]
		nx := neighbor[1]
		if ny < 0 || ny >= len(sol.grid) || nx < 0 || nx >= len(sol.grid[ny]) {
			continue
		}
		if sol.grid[ny][nx] {
			adj++
			if adj >= 4 {
				return false
			}
		}
	}
	return true
}
