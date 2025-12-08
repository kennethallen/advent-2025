package advent_2025

import (
	"iter"
	"log"
	"strings"

	"github.com/sbromberger/bitvec"
)

type Day07 struct {
	splits uint
}

func (sol *Day07) Process(input string) {
	next, stop := iter.Pull(strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n"))
	defer stop()
	first_line, ok := next()
	if !ok {
		log.Fatal()
	}
	width := uint64(len(first_line))
	row := bitvec.New(width)
	for i, c := range first_line {
		if c == 'S' {
			err := row.Set(uint64(i))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	for line, ok := next(); ok; line, ok = next() {
		next_row := bitvec.New(width)

		for i, c := range line {
			cell, err := row.Get(uint64(i))
			if err != nil {
				log.Fatal(err)
			}
			if cell {
				if c == '^' {
					sol.splits++
					if i-1 >= 0 {
						err := next_row.Set(uint64(i - 1))
						if err != nil {
							log.Fatal(err)
						}
					}
					if i+1 < len(line) {
						err := next_row.Set(uint64(i + 1))
						if err != nil {
							log.Fatal(err)
						}
					}
				} else {
					err := next_row.Set(uint64(i))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}

		row = next_row
	}
}

func (sol *Day07) Part1() uint { return sol.splits }
func (sol *Day07) Part2() uint { return 0 }
