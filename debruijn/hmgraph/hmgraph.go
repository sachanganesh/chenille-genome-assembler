package hmgraph

import (
	"velour/debruijn"
)

// ===================================
// HMGraph
// ===================================

type HMGraph struct {
	lookup			map[debruijn.Kmer]int
	nodes			[]*debruijn.GraphNode
	newNode			debruijn.NodeGenerator
}

func NewGraph(newNode debruijn.NodeGenerator) debruijn.Graph {
	var graph debruijn.Graph = &HMGraph{make(map[debruijn.Kmer]int), make([]*debruijn.GraphNode, 0, 3000000), newNode}
	return graph
}

// ===================================
// HMGraph Functions
// ===================================

func (graph *HMGraph) Len() int {
	return len(graph.nodes)
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

func (graph *HMGraph) NewNode(kmer debruijn.Kmer) debruijn.GraphNode {
	return graph.newNode(kmer)
}

func (graph *HMGraph) GetID(kmer debruijn.Kmer) (int, bool) {
	kmer_id, ok := graph.lookup[kmer]
	return kmer_id, ok
}

func (graph *HMGraph) SetID(kmer debruijn.Kmer, kmer_id int) {
	graph.lookup[kmer] = kmer_id
}

func (graph *HMGraph) GetNode(kmer debruijn.Kmer) (int, debruijn.GraphNode, bool) {
	var node debruijn.GraphNode

	if kmer_id, ok := graph.GetID(kmer); ok {
		node = *graph.nodes[kmer_id]
		return kmer_id, node, ok
	} else {
		return -1, node, ok
	}
}

func (graph *HMGraph) GetNodeByID(kmer_id int) debruijn.GraphNode {
	return *graph.nodes[kmer_id]
}

func (graph *HMGraph) SetNode(kmer debruijn.Kmer, node debruijn.GraphNode) int {
	kmer_id := graph.Len()
	graph.nodes = append(graph.nodes, &node)
	graph.SetID(kmer, kmer_id)

	return kmer_id
}

func (graph *HMGraph) ConnectNodeToGraph(kmer debruijn.Kmer, kmer_id int, node debruijn.GraphNode) {
	nts := [4]byte{'A', 'C', 'G', 'T'}

	for i, nt := range nts {
		prec_kmer := kmer.GeneratePredecessor(nt)

		if _, prec_node, ok := graph.GetNode(prec_kmer); ok {
			node.AddPredecessor(i)
			prec_node.AddSuccessor(kmer.GetLastNucleotide())
			break
		}
	}
}

func (graph *HMGraph) AddNode(kmer debruijn.Kmer) int {
	var kmer_id int
	var node debruijn.GraphNode
	var ok bool

	if kmer_id, node, ok = graph.GetNode(kmer); ok {
		node.IncrementFrequency()
	} else {
		node = graph.newNode(kmer)
		kmer_id = graph.SetNode(kmer, node)
		graph.ConnectNodeToGraph(kmer, kmer_id, node)
	}

	return kmer_id
}

func (graph *HMGraph) AddNodes(kmers []debruijn.Kmer) []int {
	ids := make([]int, 0)

	for _, kmer := range kmers {
		ids = append(ids, graph.AddNode(kmer))
	}

	return ids
}
