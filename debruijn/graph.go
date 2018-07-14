package debruijn

type Graph interface {
	Len()											int
	GetFrequencies()								[]int
	GetNumNodesSeen()								int
	NewNode(string)									GraphNode
	GetNode(string)							(int, GraphNode, bool)
	SetNode(string, GraphNode)						int
	ConnectNodeToGraph(string, int, GraphNode)
	AddNode(string)									int
	AddNodes([]string)								[]int
}

type GraphNode interface {
	GetKmer()				string
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

type NodeGenerator func(string) GraphNode
