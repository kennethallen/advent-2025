package advent_2025

import (
	"container/heap"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Day08 struct {
	prod_of_sizes_of_three_largest_circuits_after_benchmark uint64
	prod_of_x_coordinates_of_last_two_junctions_to_connect  uint64
}

type JunctionId int
type CircuitId int

type Slot struct {
	junction Junction
	circuit  []JunctionId
}

type Junction struct {
	coord      [3]uint64
	circuit_id CircuitId
}

type JunctionPair struct {
	dist_sq   uint64
	junctions [2]JunctionId
}

type JunctionPairHeap []JunctionPair

func (pairs JunctionPairHeap) Len() int           { return len(pairs) }
func (pairs JunctionPairHeap) Less(i, j int) bool { return pairs[i].dist_sq < pairs[j].dist_sq }
func (pairs JunctionPairHeap) Swap(i, j int)      { pairs[i], pairs[j] = pairs[j], pairs[i] }
func (pairs *JunctionPairHeap) Push(x any)        { *pairs = append(*pairs, x.(JunctionPair)) }
func (pairs *JunctionPairHeap) Pop() any {
	val := (*pairs)[len(*pairs)-1]
	*pairs = (*pairs)[:len(*pairs)-1]
	return val
}

func (sol *Day08) Process(input string) {
	var slots []Slot
	for line := range strings.SplitSeq(strings.TrimSuffix(input, "\n"), "\n") {
		coords := strings.Split(line, ",")
		x, err := strconv.ParseUint(coords[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseUint(coords[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.ParseUint(coords[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		slots = append(slots, Slot{
			junction: Junction{
				coord:      [3]uint64{x, y, z},
				circuit_id: CircuitId(len(slots)),
			},
			circuit: []JunctionId{JunctionId(len(slots))},
		})
	}

	pairs := JunctionPairHeap(make([]JunctionPair, 0, (len(slots)*len(slots)-len(slots))/2))
	for i, slot_i := range slots {
		junc_i := slot_i.junction
		for j := i + 1; j < len(slots); j++ {
			junc_j := slots[j].junction
			dx := junc_i.coord[0] - junc_j.coord[0]
			dy := junc_i.coord[1] - junc_j.coord[1]
			dz := junc_i.coord[2] - junc_j.coord[2]
			pairs = append(pairs, JunctionPair{
				junctions: [2]JunctionId{JunctionId(i), JunctionId(j)},
				dist_sq:   dx*dx + dy*dy + dz*dz,
			})
		}
	}
	heap.Init(&pairs)
	circuits := len(slots)
	for n := 0; ; n++ {
		if len(slots) == 20 && n == 10 || len(slots) != 20 && n == 1_000 { // Example vs real
			circuit_sizes := make([]int, 0, len(slots))
			for _, slot := range slots {
				circuit_sizes = append(circuit_sizes, len(slot.circuit))
			}
			slices.Sort(circuit_sizes)
			sol.prod_of_sizes_of_three_largest_circuits_after_benchmark = 1
			for _, circuit_size := range circuit_sizes[len(circuit_sizes)-3:] {
				sol.prod_of_sizes_of_three_largest_circuits_after_benchmark *= uint64(circuit_size)
			}
		}

		closest_pair := heap.Pop(&pairs).(JunctionPair)
		j0 := closest_pair.junctions[0]
		j0_circ := slots[j0].junction.circuit_id
		j1 := closest_pair.junctions[1]
		j1_circ := slots[j1].junction.circuit_id
		if j0_circ == j1_circ { // Already connected
			continue
		}
		if len(slots[j0_circ].circuit) < len(slots[j1_circ].circuit) { // Merge into larger
			j0_circ, j1_circ = j1_circ, j0_circ
		}
		slots[j0_circ].circuit = append(slots[j0_circ].circuit, slots[j1_circ].circuit...)
		for _, j := range slots[j1_circ].circuit {
			slots[j].junction.circuit_id = j0_circ
		}
		slots[j1_circ].circuit = nil

		circuits--
		if circuits == 1 {
			sol.prod_of_x_coordinates_of_last_two_junctions_to_connect = slots[j0].junction.coord[0] * slots[j1].junction.coord[0]
			break
		}
	}
}

func (sol *Day08) Part1() uint64 { return sol.prod_of_sizes_of_three_largest_circuits_after_benchmark }
func (sol *Day08) Part2() uint64 { return sol.prod_of_x_coordinates_of_last_two_junctions_to_connect }
