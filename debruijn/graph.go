package debruijn

import (
	"strings"
)

type Graph interface {
	Len()											int
	GetID(string)									(int, bool)
	SetID(string, int)
	GetNodeByKey(string)							(int, GraphNode, bool)
	GetNodeByID(int)								(GraphNode)
	SetNode(GraphNode)								int
	ConnectNodeToGraph(string, int, GraphNode)
	AddNode(string)									int
	AddNodes([]string)								[]int
}

// ===================================
// DBGraph
// ===================================

type DBGraph struct {
	lookup			map[string]int
	nodes			[]*GraphNode
}

func NewGraph() *DBGraph {
	return &DBGraph{make(map[string]int), make([]*GraphNode, 0, 3000000)}
}

// ===================================
// DBGraph Functions
// ===================================

func (graph *DBGraph) Len() int {
	return len(graph.nodes)
}

func (graph *DBGraph) GetID(kmer string) (int, bool) {
	kmer_id, ok := graph.lookup[kmer]
	return kmer_id, ok
}

func (graph *DBGraph) SetID(kmer string, kmer_id int) {
	graph.lookup[kmer] = kmer_id
}

func (graph *DBGraph) GetNodeByKey(kmer string) (int, GraphNode, bool) {
	var kmer_entry GraphNode

	if kmer_id, ok := graph.GetID(kmer); ok {
		kmer_entry = *graph.nodes[kmer_id]
		return kmer_id, kmer_entry, ok
	} else {
		return -1, kmer_entry, ok
	}
}

func (graph *DBGraph) GetNodeByID(kmer_id int) GraphNode {
	return *graph.nodes[kmer_id]
}

func (graph *DBGraph) SetNode(kmer_entry GraphNode) int {
	kmer_id := graph.Len()
	graph.nodes = append(graph.nodes, &kmer_entry)

	return kmer_id
}

func (graph *DBGraph) ConnectNodeToGraph(kmer string, kmer_id int, kmer_entry GraphNode) {
	nts := [...]string{"A", "C", "G", "T"}

	base := kmer[1 : len(kmer)]
	for _, nt := range nts {
		var prec_kmer_buf strings.Builder
		prec_kmer_buf.WriteString(nt)
		prec_kmer_buf.WriteString(base)
		prec_kmer := prec_kmer_buf.String()

		if prec_id, prec_kmer_entry, ok := graph.GetNodeByKey(prec_kmer); ok {
			kmer_entry.AddPredecessor(prec_id)
			prec_kmer_entry.AddSuccessor(kmer_id)
		}
	}
}

func (graph *DBGraph) AddNode(kmer string) int {
	var kmer_id int
	var kmer_entry GraphNode
	var ok bool

	if kmer_id, kmer_entry, ok = graph.GetNodeByKey(kmer); ok {
		kmer_entry.IncrementFrequency()
	} else {
		kmer_entry = NewNode(kmer)
		kmer_id = graph.SetNode(kmer_entry)
		graph.ConnectNodeToGraph(kmer, kmer_id, kmer_entry)
		graph.SetID(kmer, kmer_id)
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
