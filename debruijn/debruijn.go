package debruijn

import (
	"strings"
)

// ===================================
// KmerEntry Utilities
// ===================================

type KmerEntry struct {
	Kmer			string
	Frequency		int
	Predecessors	[]int
	Successors		[]int
}

func NewKmerEntry(kmer string) KmerEntry {
	return KmerEntry{kmer, 0, make([]int, 0), make([]int, 0)}
}

// ===================================
// DBGraph Utilities
// ===================================

type DBGraph struct {
	NumNodes		int
	Lookup			map[string]int
	Nodes			[]*KmerEntry
}

func NewDBGraph() DBGraph {
	return DBGraph{0, make(map[string]int), make([]*KmerEntry, 0)}
}

func (graph *DBGraph) IncludesKmer(kmer string) bool {
	_, ok := graph.Lookup[kmer]
	return ok
}

func (graph *DBGraph) IsInPredecessor(kmer_id int, needle int) bool {
	if kmer_id < len(graph.Nodes) {
		kmer_entry := *graph.Nodes[kmer_id]

		for i := range kmer_entry.Predecessors {
			if kmer_entry.Predecessors[i] == needle {
				return true
			}
		}
	}

	return false
}

func (graph *DBGraph) IsInSuccessor(kmer_id int, needle int) bool {
	if kmer_id < len(graph.Nodes) {
		kmer_entry := *graph.Nodes[kmer_id]

		for i := range kmer_entry.Successors {
			if kmer_entry.Successors[i] == needle {
				return true
			}
		}
	}

	return false
}

func (graph *DBGraph) precedingKmers(kmer string, prec_kmer_id int) []int {
	nts := [...]string{"A", "C", "G", "T"}
	prec_kmer_ids := make([]int, 1)

	if prec_kmer_id != -1 {
		prec_kmer_ids = append(prec_kmer_ids, prec_kmer_id)
	}

	base := kmer[1 : len(kmer)]
	for i := range nts {
		var prec_kmer_buf strings.Builder
		prec_kmer_buf.WriteString(nts[i])
		prec_kmer_buf.WriteString(base)
		prec_kmer := prec_kmer_buf.String()

		if prec_kmer_id != -1 && graph.Nodes[prec_kmer_id].Kmer != prec_kmer && graph.IncludesKmer(prec_kmer) {
			prec_kmer_ids = append(prec_kmer_ids, graph.Lookup[prec_kmer])
		} else if prec_kmer_id == -1 && graph.IncludesKmer(prec_kmer) {
			prec_kmer_ids = append(prec_kmer_ids, graph.Lookup[prec_kmer])
		}
	}

	return prec_kmer_ids
}

func (graph *DBGraph) AddNode(kmer string, prec_kmer_id int) int {
	prec_kmers := graph.precedingKmers(kmer, prec_kmer_id)

	var kmer_entry KmerEntry
	var kmer_id int

	if graph.IncludesKmer(kmer) {
		kmer_id = graph.Lookup[kmer]
		kmer_entry = *graph.Nodes[kmer_id]
	} else {
		kmer_id = graph.NumNodes
		graph.NumNodes += 1

		graph.Lookup[kmer] = kmer_id

		kmer_entry = NewKmerEntry(kmer)
		graph.Nodes = append(graph.Nodes, &kmer_entry)
	}

	kmer_entry.Frequency += 1

	for i := range prec_kmers {
		prec_kmer_id := prec_kmers[i]

		if !graph.IsInPredecessor(kmer_id, prec_kmer_id) {
			kmer_entry.Predecessors = append(kmer_entry.Predecessors, prec_kmer_id)
		}

		if !graph.IsInSuccessor(prec_kmer_id, kmer_id) {
			prec_kmer_entry := *graph.Nodes[prec_kmer_id]
			prec_kmer_entry.Successors = append(prec_kmer_entry.Successors, kmer_id)
		}
	}

	return kmer_id
}

func (graph *DBGraph) AddNodes(kmers []string) {
	prec_kmer_id := -1

	for i := range kmers {
		prec_kmer_id = graph.AddNode(kmers[i], prec_kmer_id)
	}
}
