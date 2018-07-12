package debruijn_sync

import (
	"sync"
	"velour/debruijn"
)

// ===================================
// DBGraph
// ===================================

type DBGraph struct {
	data			debruijn.Graph
	mu				sync.RWMutex
}

func NewGraph() *DBGraph {
	return &DBGraph{debruijn.NewGraph(), sync.RWMutex{}}
}

func NewGraphFrom(graph debruijn.Graph) *DBGraph {
	return &DBGraph{graph, sync.RWMutex{}}
}

// ===================================
// DBGraph Functions
// ===================================

func (graph *DBGraph) Len() int {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.Len()
}

func (graph *DBGraph) GetID(kmer string) (int, bool) {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.GetID(kmer)
}

func (graph *DBGraph) SetID(kmer string, kmer_id int) {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	graph.data.SetID(kmer, kmer_id)
}

func (graph *DBGraph) GetNodeByKey(kmer string) (int, debruijn.GraphNode, bool) {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.GetNodeByKey(kmer)
}

func (graph *DBGraph) GetNodeByID(kmer_id int) debruijn.GraphNode {
	graph.mu.RLock()
	defer graph.mu.RUnlock()

	return graph.data.GetNodeByID(kmer_id)
}

func (graph *DBGraph) SetNode(kmer_entry debruijn.GraphNode) int {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	return graph.data.SetNode(kmer_entry)
}

func (graph *DBGraph) ConnectNodeToGraph(kmer string, kmer_id int, kmer_entry debruijn.GraphNode) {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	graph.data.ConnectNodeToGraph(kmer, kmer_id, kmer_entry)
}

func (graph *DBGraph) addKmerEntry(kmer string, kmer_entry debruijn.GraphNode) int {
	graph.mu.Lock()
	defer graph.mu.Unlock()

	var kmer_id int

	if og_id, og_entry, ok := graph.data.GetNodeByKey(kmer); ok {
		og_entry.Merge(kmer_entry)
		kmer_entry = og_entry
		kmer_id = og_id
	} else {
		kmer_id = graph.data.SetNode(kmer_entry)
		graph.data.SetID(kmer, kmer_id)
	}

	graph.data.ConnectNodeToGraph(kmer, kmer_id, kmer_entry)

	return kmer_id
}

func (graph *DBGraph) AddNode(kmer string) int {
	var kmer_id int
	var kmer_entry debruijn.GraphNode
	var ok bool

	if kmer_id, kmer_entry, ok = graph.GetNodeByKey(kmer); ok {
		kmer_entry.IncrementFrequency()
	} else {
		kmer_entry = NewNode(kmer)
		kmer_id = graph.addKmerEntry(kmer, kmer_entry)
	}

	return kmer_id
}

func (graph *DBGraph) AddNodes(kmers []string) []int {
	ids := make([]int, 0)

	for _, kmer := range kmers {
		ids = append(ids, graph.AddNode(kmer))
	}

	return ids
}
