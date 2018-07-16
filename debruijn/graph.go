package debruijn

// ===================================
// Graph
// ===================================

type Graph interface {
	Len()											int
	GetFrequencies()								[]int
	GetNumNodesSeen()								int
	NewNode(Kmer)									GraphNode
	GetNode(Kmer)									(int, GraphNode, bool)
	SetNode(Kmer, GraphNode)						int
	ConnectNodeToGraph(Kmer, int, GraphNode)
	AddNode(Kmer)									int
	AddNodes([]Kmer)								[]int
}

type GraphNode interface {
	GetKmer()				Kmer
	GetFrequency()			int
	SetFrequency(int)
	IncrementFrequency()
	GetPredecessors()		[]int
	IsAPredecessor(int)		bool
	AddPredecessor(int)
	GetSuccessors()			[]int
	IsASuccessor(int)		bool
	AddSuccessor(int)
	Merge(GraphNode)
}

type NodeGenerator func(Kmer) GraphNode
