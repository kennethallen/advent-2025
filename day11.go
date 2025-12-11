package advent_2025

import (
	"log"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/sbromberger/bitvec"
)

type Day11 struct {
	names map[string]int
	devs  [][]int
}

type DeviceParse struct {
	Id    string   `parser:"@Ident ':'"`
	Conns []string `parser:"@Ident*"`
}

func (sol *Day11) Process(input string) {
	parser, err := participle.Build[DeviceParse]()
	if err != nil {
		log.Fatal(err)
	}
	sol.names = make(map[string]int)
	translate := func(name string) int {
		id, ok := sol.names[name]
		if !ok {
			id = len(sol.names)
			sol.names[name] = id
		}
		return id
	}
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n") {
		ast, err := parser.ParseString("", line)
		if err != nil {
			log.Fatal(err)
		}
		id := translate(ast.Id)
		conns := make([]int, 0, len(ast.Conns))
		for _, conn := range ast.Conns {
			conns = append(conns, translate(conn))
		}
		if id >= len(sol.devs) {
			// Double size, or more if necessary
			sol.devs = append(sol.devs, make([][]int, max(id+1-len(sol.devs), len(sol.devs)))...)
		}
		sol.devs[id] = conns
	}
}

func (sol *Day11) recurse11(curs int, cache []uint64, fill_vec bitvec.BitVec) uint64 {
	filled, err := fill_vec.Get(uint64(curs))
	if err != nil {
		log.Fatal(err)
	}
	if filled {
		return cache[curs]
	}

	var val uint64
	for _, next := range sol.devs[curs] {
		val += sol.recurse11(next, cache, fill_vec)
	}
	cache[curs] = val
	err = fill_vec.Set(uint64(curs))
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func (sol *Day11) Part1() uint64 {
	cache := make([]uint64, len(sol.devs))
	fill_vec := bitvec.New(uint64(len(sol.devs)))

	out := sol.names["out"]
	cache[out] = 1
	err := fill_vec.Set(uint64(out))
	if err != nil {
		log.Fatal(err)
	}

	return sol.recurse11(sol.names["you"], cache, fill_vec)
}

type part_2_paths struct {
	neither, dac, fft, both uint64
}

func (sol *Day11) Part2() uint64 {
	cache := make([]part_2_paths, len(sol.devs))
	fill_vec := bitvec.New(uint64(len(sol.devs)))

	out := sol.names["out"]
	cache[out] = part_2_paths{neither: 1}
	err := fill_vec.Set(uint64(out))
	if err != nil {
		log.Fatal(err)
	}

	return sol.recurse11_2(sol.names["svr"], sol.names["dac"], sol.names["fft"], cache, fill_vec).both
}

func (sol *Day11) recurse11_2(curs, dac, fft int, cache []part_2_paths, fill_vec bitvec.BitVec) part_2_paths {
	filled, err := fill_vec.Get(uint64(curs))
	if err != nil {
		log.Fatal(err)
	}
	if filled {
		return cache[curs]
	}

	var val part_2_paths
	for _, next := range sol.devs[curs] {
		next_val := sol.recurse11_2(next, dac, fft, cache, fill_vec)
		val.neither += next_val.neither
		val.dac += next_val.dac
		val.fft += next_val.fft
		val.both += next_val.both
	}
	switch curs {
	case dac:
		val = part_2_paths{
			dac:  val.dac + val.neither,
			both: val.both + val.fft,
		}
	case fft:
		val = part_2_paths{
			fft:  val.fft + val.neither,
			both: val.both + val.dac,
		}
	}
	cache[curs] = val
	err = fill_vec.Set(uint64(curs))
	if err != nil {
		log.Fatal(err)
	}
	return val
}
