package advent_2025

import (
	"container/heap"
	"log"
	"slices"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Day08 struct {
	prod_of_sizes_of_three_largest_circuits uint64
}

type JunctionId int
type CircuitId int

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
	var junctions []Junction
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
		junctions = append(junctions, Junction{coord: [3]uint64{x, y, z}})
	}

	pairs := JunctionPairHeap(make([]JunctionPair, 0, (len(junctions)*len(junctions)-len(junctions))/2))
	for i, junc_i := range junctions {
		for j := i + 1; j < len(junctions); j++ {
			junc_j := junctions[j]
			dx := junc_i.coord[0] - junc_j.coord[0]
			dy := junc_i.coord[1] - junc_j.coord[1]
			dz := junc_i.coord[2] - junc_j.coord[2]
			pairs = append(pairs, JunctionPair{junctions: [2]JunctionId{JunctionId(i), JunctionId(j)}, dist_sq: dx*dx + dy*dy + dz*dz})
		}
	}
	heap.Init(&pairs)
	var circuits []mapset.Set[JunctionId]
	pairs_to_link := 1_000
	if len(junctions) == 20 {
		pairs_to_link = 10
	}
	for range pairs_to_link {
		closest_pair := heap.Pop(&pairs).(JunctionPair)
		if junctions[closest_pair.junctions[0]].circuit_id == 0 { // First is not in a circuit
			if junctions[closest_pair.junctions[1]].circuit_id == 0 { // Nor is second
				// New circuit!
				new_circuit_id := len(circuits) + 1
				log.Printf("New circuit %v for junctions %v", new_circuit_id, closest_pair.junctions)
				for i := range closest_pair.junctions {
					junctions[closest_pair.junctions[i]].circuit_id = CircuitId(new_circuit_id)
				}
				circuits = append(circuits, mapset.NewThreadUnsafeSet(closest_pair.junctions[:]...))
			} else { // Second is
				// First joins second's circuit
				circuit_id := junctions[closest_pair.junctions[1]].circuit_id
				log.Printf("First %v joining circuit %v", closest_pair.junctions[0], circuit_id)
				junctions[closest_pair.junctions[0]].circuit_id = circuit_id
				circuits[circuit_id-1].Add(closest_pair.junctions[0])
			}
		} else { // First is in a circuit
			circuit_id := junctions[closest_pair.junctions[0]].circuit_id
			if junctions[closest_pair.junctions[1]].circuit_id == 0 { // Second isn't
				// Second joins first's circuit
				log.Printf("Second %v joining circuit %v", closest_pair.junctions[1], circuit_id)
				junctions[closest_pair.junctions[1]].circuit_id = circuit_id
				circuits[circuit_id-1].Add(closest_pair.junctions[1])
			} else if loser := junctions[closest_pair.junctions[1]].circuit_id; circuit_id != loser { // Second is
				// Merging circuits
				log.Printf("Merging members of circuit %v into %v", loser, circuit_id)
				for junction_id := range circuits[loser-1].Iter() {
					junctions[junction_id].circuit_id = circuit_id
					circuits[circuit_id-1].Add(junction_id)
				}
				circuits[loser-1] = nil
			}
		}

		// Check invariants
		for junction_id, junction := range junctions {
			if junction.circuit_id != 0 {
				if circuits[junction.circuit_id-1] == nil {
					log.Fatalf("Junction %v circuit %v nil", junction_id, junction.circuit_id)
				}
				if !circuits[junction.circuit_id-1].ContainsOne(JunctionId(junction_id)) {
					log.Fatalf("Junction %v not in circuit %v", junction_id, junction.circuit_id)
				}
			}
		}
		for circuit_id_minus_one, circuit := range circuits {
			if circuit == nil {
				continue
			}
			if circuit.Cardinality() == 0 {
				log.Fatalf("Circuit %v empty", circuit_id_minus_one+1)
			}
			for junction_id := range circuit.Iter() {
				if junctions[junction_id].circuit_id != CircuitId(circuit_id_minus_one+1) {
					log.Fatalf("Circuit %v not set on member %v", circuit_id_minus_one+1, junction_id)
				}
			}
		}
	}

	slices.SortFunc(circuits, func(a, b mapset.Set[JunctionId]) int {
		var a_card, b_card int
		if a != nil {
			a_card = a.Cardinality()
		}
		if b != nil {
			b_card = b.Cardinality()
		}
		return a_card - b_card
	})
	log.Println(circuits)
	sol.prod_of_sizes_of_three_largest_circuits = uint64(circuits[len(circuits)-1].Cardinality()) * uint64(circuits[len(circuits)-3].Cardinality()) * uint64(circuits[len(circuits)-2].Cardinality())
}

func (sol *Day08) Part1() uint64 { return sol.prod_of_sizes_of_three_largest_circuits }
func (sol *Day08) Part2() uint64 { return 0 }
