package hmgraph

import (
	"strings"
	"velour/debruijn"
)

// ===================================
// HMGraph
// ===================================

type HMGraph struct {
	lookup			map[string]int
	nodes			[]*debruijn.GraphNode
	newNode			debruijn.NodeGenerator
}

func NewGraph(newNode debruijn.NodeGenerator) debruijn.Graph {
	var graph debruijn.Graph = &HMGraph{make(map[string]int), make([]*debruijn.GraphNode, 0, 3000000), newNode}
	return graph
}

// ===================================
// HMGraph Functions
// ===================================

func (graph *HMGraph) Len() int {
	return len(graph.nodes)
}

func (graph *HMGraph) NewNode(kmer string) debruijn.GraphNode {
	return graph.newNode(kmer)
}

func (graph *HMGraph) GetID(kmer string) (int, bool) {
	kmer_id, ok := graph.lookup[kmer]
	return kmer_id, ok
}

func (graph *HMGraph) SetID(kmer string, kmer_id int) {
	graph.lookup[kmer] = kmer_id
}

func (graph *HMGraph) GetNode(kmer string) (int, debruijn.GraphNode, bool) {
	var kmer_entry debruijn.GraphNode

	if kmer_id, ok := graph.GetID(kmer); ok {
		kmer_entry = *graph.nodes[kmer_id]
		return kmer_id, kmer_entry, ok
	} else {
		return -1, kmer_entry, ok
	}
}

func (graph *HMGraph) GetNodeByID(kmer_id int) debruijn.GraphNode {
	return *graph.nodes[kmer_id]
}

func (graph *HMGraph) GetFrequencies() []int {
	freqs := make([]int, graph.Len())

	for i := range freqs {
		freqs[i] = graph.GetNodeByID(i).GetFrequency()
	}

	return freqs
}

func (graph *HMGraph) GetNumNodesSeen() int {
	num_seen := 0

	for _, freq := range graph.GetFrequencies() {
		num_seen += freq
	}

	return num_seen
}

func (graph *HMGraph) SetNode(kmer string, kmer_entry debruijn.GraphNode) int {
	kmer_id := graph.Len()
	graph.nodes = append(graph.nodes, &kmer_entry)
	graph.SetID(kmer, kmer_id)

	return kmer_id
}

func (graph *HMGraph) ConnectNodeToGraph(kmer string, kmer_id int, kmer_entry debruijn.GraphNode) {
	nts := [...]string{"A", "C", "G", "T"}

	base := kmer[1 : len(kmer)]
	for _, nt := range nts {
		var prec_kmer_buf strings.Builder
		prec_kmer_buf.WriteString(nt)
		prec_kmer_buf.WriteString(base)
		prec_kmer := prec_kmer_buf.String()

		if prec_id, prec_kmer_entry, ok := graph.GetNode(prec_kmer); ok {
			kmer_entry.AddPredecessor(prec_id)
			prec_kmer_entry.AddSuccessor(kmer_id)
		}
	}
}

func (graph *HMGraph) AddNode(kmer string) int {
	var kmer_id int
	var kmer_entry debruijn.GraphNode
	var ok bool

	if kmer_id, kmer_entry, ok = graph.GetNode(kmer); ok {
		kmer_entry.IncrementFrequency()
	} else {
		kmer_entry = graph.NewNode(kmer)
		kmer_id = graph.SetNode(kmer, kmer_entry)
		graph.ConnectNodeToGraph(kmer, kmer_id, kmer_entry)
	}

	return kmer_id
}

func (graph *HMGraph) AddNodes(kmers []string) []int {
	ids := make([]int, 0)

	for _, kmer := range kmers {
		ids = append(ids, graph.AddNode(kmer))
	}

	return ids
}
