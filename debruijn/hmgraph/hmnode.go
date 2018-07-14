package hmgraph

import (
	"velour/debruijn"
)

// ===================================
// HMNode
// ===================================

type HMNode struct {
	kmer			string
	frequency		int
	predecessors	[]int
	successors		[]int
}

func NewNode(kmer string) debruijn.GraphNode {
	var node debruijn.GraphNode = &HMNode{kmer, 1, make([]int, 0, 2), make([]int, 0, 2)}
	return node
}

// ===================================
// HMNode Functions
// ===================================

func (node *HMNode) GetKmer() string {
	return node.kmer
}

func (node *HMNode) GetFrequency() int {
	return node.frequency
}

func (node *HMNode) SetFrequency(new_frequency int) {
	node.frequency = new_frequency
}

func (node *HMNode) IncrementFrequency() {
	node.SetFrequency(node.GetFrequency() + 1)
}

func (node *HMNode) GetPredecessors() []int {
	return node.predecessors
}

func (node *HMNode) IsAPredecessor(needle int) bool {
	for _, pred := range node.predecessors {
		if pred == needle {
			return true
		}
	}

	return false
}

func (node *HMNode) AddPredecessor(pred_id int) {
	if !node.IsAPredecessor(pred_id) {
		node.predecessors = append(node.predecessors, pred_id)
	}
}

func (node *HMNode) GetSuccessors() []int {
	return node.successors
}

func (node *HMNode) IsASuccessor(needle int) bool {
	for _, succ := range node.successors {
		if succ == needle {
			return true
		}
	}

	return false
}

func (node *HMNode) AddSuccessor(succ_id int) {
	if !node.IsASuccessor(succ_id) {
		node.successors = append(node.successors, succ_id)
	}
}

func (node *HMNode) Merge(other_entry debruijn.GraphNode) {
	if node.GetKmer() == other_entry.GetKmer() {
		node.SetFrequency(node.GetFrequency() + other_entry.GetFrequency())

		for _, pred_id := range other_entry.GetPredecessors() {
			node.AddPredecessor(pred_id)
		}

		for _, succ_id := range other_entry.GetSuccessors() {
			node.AddSuccessor(succ_id)
		}
	}
}
