package advent_2025

import (
	"iter"
	"log"
	"strings"

	"github.com/sbromberger/bitvec"
)

type Day07 struct {
	width  uint64
	starts []uint64
	rows   []bitvec.BitVec
}

func (sol *Day07) Process(input string) {
	next, stop := iter.Pull(strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n"))
	defer stop()
	first_line, ok := next()
	if !ok {
		log.Fatal()
	}
	sol.width = uint64(len(first_line))
	for i, c := range first_line {
		if c == 'S' {
			sol.starts = append(sol.starts, uint64(i))
		}
	}

	for line, ok := next(); ok; line, ok = next() {
		row := bitvec.New(sol.width)
		for i, c := range line {
			if c == '^' {
				err := row.Set(uint64(i))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		sol.rows = append(sol.rows, row)
	}
}

func (sol *Day07) Part1() uint {
	var splits uint
	tachs := bitvec.New(sol.width)
	for _, start := range sol.starts {
		err := tachs.Set(start)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, row := range sol.rows {
		next_tachs := bitvec.New(sol.width)
		for i := range sol.width {
			tach, err := tachs.Get(i)
			if err != nil {
				log.Fatal(err)
			}
			if !tach {
				continue
			}
			splitter, err := row.Get(i)
			if err != nil {
				log.Fatal(err)
			}
			if splitter {
				splits++
				if i > 0 {
					err := next_tachs.Set(i - 1)
					if err != nil {
						log.Fatal(err)
					}
				}
				if i+1 < sol.width {
					err := next_tachs.Set(i + 1)
					if err != nil {
						log.Fatal(err)
					}
				}
			} else {
				err := next_tachs.Set(i)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		tachs = next_tachs
	}
	return splits
}

type Day07Part2 struct {
	base  *Day07
	cache map[[2]uint64]uint
}

func (sol *Day07) Part2() uint {
	part2 := Day07Part2{base: sol, cache: make(map[[2]uint64]uint)}
	var universes uint
	for _, start := range sol.starts {
		universes += part2.recurse(0, start)
	}
	return universes
}

func (sol *Day07Part2) recurse(depth, pos uint64) uint {
	key := [2]uint64{depth, pos}
	val, cached := sol.cache[key]
	if !cached {
		if depth >= uint64(len(sol.base.rows)) {
			val = 1
		} else {
			splitter, err := sol.base.rows[depth].Get(pos)
			if err != nil {
				log.Fatal(err)
			}
			if splitter {
				if pos > 0 {
					val = sol.recurse(depth+1, pos-1)
				}
				if pos+1 < uint64(len(sol.base.rows)) {
					val += sol.recurse(depth+1, pos+1)
				}
			} else {
				val = sol.recurse(depth+1, pos)
			}
		}
		sol.cache[key] = val
	}
	return val
}
