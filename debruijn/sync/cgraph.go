package sync

import (
	"sync"
	"velour/debruijn"
)

// ===================================
// CGraph
// ===================================

type CGraph struct {
	data			debruijn.Graph
	mu				sync.RWMutex
}

func NewGraph(graph debruijn.Graph) debruijn.Graph {
	var graph debruijn.Graph = &CGraph{graph, sync.RWMutex{}}
	return graph
}

// ===================================
// CGraph Functions
// ===================================

func (graph *CGraph) Len() int {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.Len()
}

func (graph *CGraph) GetFrequencies() []int {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.GetFrequencies()
}

func (graph *CGraph) GetNumNodesSeen() int {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.GetNumNodesSeen()
}

func (graph *CGraph) NewNode(kmer string) debruijn.GraphNode {
	return graph.data.NewNode(kmer)
}

func (graph *CGraph) GetNode(kmer string) (int, debruijn.GraphNode, bool) {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.GetNode(kmer)
}

func (graph *CGraph) SetNode(kmer string, node debruijn.GraphNode) int {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	return graph.data.SetNode(kmer, node)
}

func (graph *CGraph) ConnectNodeToGraph(kmer string, kmer_id int, node debruijn.GraphNode) {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	graph.data.ConnectNodeToGraph(kmer, kmer_id, node)
}

func (graph *CGraph) addKmerEntry(kmer string, node debruijn.GraphNode) int {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	var kmer_id int

	if og_id, og_entry, ok := graph.data.GetNode(kmer); ok {
		og_entry.Merge(node)
		node = og_entry
		kmer_id = og_id
	} else {
		kmer_id = graph.data.SetNode(kmer, node)
	}

	graph.data.ConnectNodeToGraph(kmer, kmer_id, node)

	return kmer_id
}

func (graph *CGraph) AddNode(kmer string) int {
	var kmer_id int
	var node debruijn.GraphNode
	var ok bool

	if kmer_id, node, ok = graph.GetNode(kmer); ok {
		node.IncrementFrequency()
	} else {
		node = graph.NewNode(kmer)
		kmer_id = graph.addKmerEntry(kmer, node)
	}

	return kmer_id
}

func (graph *CGraph) AddNodes(kmers []string) []int {
	ids := make([]int, 0)

	for _, kmer := range kmers {
		ids = append(ids, graph.AddNode(kmer))
	}

	return ids
}
