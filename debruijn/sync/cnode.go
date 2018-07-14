package sync

import (
	"sync"
	"velour/debruijn"
)

// ===================================
// CNode
// ===================================

type CNode struct {
	data			debruijn.GraphNode
	mu				sync.Mutex
}

func NewNode(kmer string, newNode debruijn.NodeGenerator) *debruijn.GraphNode {
	var node debruijn.GraphNode = &CNode{newNode(kmer), sync.Mutex{}}
	return node
}

// ===================================
// CNode Functions
// ===================================

func (node *CNode) GetKmer() string {
	node.mu.Lock()
	defer node.mu.Unlock()

	return node.data.GetKmer()
}

func (node *CNode) GetFrequency() int {
	node.mu.Lock()
	defer node.mu.Unlock()

	return node.data.GetFrequency()
}

func (node *CNode) SetFrequency(new_frequency int) {
	node.mu.Lock()
	defer node.mu.Unlock()

	node.data.SetFrequency(new_frequency)
}

func (node *CNode) IncrementFrequency() {
	node.SetFrequency(node.GetFrequency() + 1)
}

func (node *CNode) GetPredecessors() []int {
	node.mu.Lock()
	defer node.mu.Unlock()

	return node.data.GetPredecessors()
}

func (node *CNode) IsAPredecessor(needle int) bool {
	node.mu.Lock()
	defer node.mu.Unlock()

	return node.data.IsAPredecessor(needle)
}

func (node *CNode) AddPredecessor(pred_id int) {
	node.mu.Lock()
	defer node.mu.Unlock()

	node.data.AddPredecessor(pred_id)
}

func (node *CNode) GetSuccessors() []int {
	node.mu.Lock()
	defer node.mu.Unlock()

	return node.data.GetSuccessors()
}

func (node *CNode) IsASuccessor(needle int) bool {
	node.mu.Lock()
	defer node.mu.Unlock()

	return node.data.IsASuccessor(needle)
}

func (node *CNode) AddSuccessor(succ_id int) {
	node.mu.Lock()
	defer node.mu.Unlock()

	node.data.AddSuccessor(succ_id)
}

func (node *CNode) Merge(other_entry debruijn.GraphNode) {
	node.mu.Lock()
	defer node.mu.Unlock()

	if node.data.GetKmer() == other_entry.GetKmer() {
		node.data.SetFrequency(node.data.GetFrequency() + other_entry.GetFrequency())

		for _, pred_id := range other_entry.GetPredecessors() {
			node.data.AddPredecessor(pred_id)
		}

		for _, succ_id := range other_entry.GetSuccessors() {
			node.data.AddSuccessor(succ_id)
		}
	}
}
