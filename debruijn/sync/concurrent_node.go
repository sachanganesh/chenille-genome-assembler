package debruijn_sync

import (
	"sync"
	"velour/debruijn"
)

// ===================================
// KmerNode
// ===================================

type KmerNode struct {
	data			debruijn.GraphNode
	mu				sync.Mutex
}

func NewNode(kmer string) *KmerNode {
	return &KmerNode{debruijn.NewNode(kmer), sync.Mutex{}}
}

// ===================================
// KmerNode Functions
// ===================================

func (kmer_entry *KmerNode) GetKmer() string {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	return kmer_entry.data.GetKmer()
}

func (kmer_entry *KmerNode) GetFrequency() int {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	return kmer_entry.data.GetFrequency()
}

func (kmer_entry *KmerNode) SetFrequency(new_frequency int) {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	kmer_entry.data.SetFrequency(new_frequency)
}

func (kmer_entry *KmerNode) IncrementFrequency() {
	kmer_entry.SetFrequency(kmer_entry.GetFrequency() + 1)
}

func (kmer_entry *KmerNode) GetPredecessors() []int {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	return kmer_entry.data.GetPredecessors()
}

func (kmer_entry *KmerNode) IsAPredecessor(needle int) bool {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	return kmer_entry.data.IsAPredecessor(needle)
}

func (kmer_entry *KmerNode) AddPredecessor(pred_id int) {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	kmer_entry.data.AddPredecessor(pred_id)
}

func (kmer_entry *KmerNode) GetSuccessors() []int {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	return kmer_entry.data.GetSuccessors()
}

func (kmer_entry *KmerNode) IsASuccessor(needle int) bool {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	return kmer_entry.data.IsASuccessor(needle)
}

func (kmer_entry *KmerNode) AddSuccessor(succ_id int) {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	kmer_entry.data.AddSuccessor(succ_id)
}

func (kmer_entry *KmerNode) Merge(other_entry debruijn.GraphNode) {
	kmer_entry.mu.Lock()
	defer kmer_entry.mu.Unlock()

	if kmer_entry.data.GetKmer() == other_entry.GetKmer() {	kmer_entry.data.SetFrequency(kmer_entry.data.GetFrequency() + other_entry.GetFrequency())

		for _, pred_id := range other_entry.GetPredecessors() {
			kmer_entry.data.AddPredecessor(pred_id)
		}

		for _, succ_id := range other_entry.GetSuccessors() {
			kmer_entry.data.AddSuccessor(succ_id)
		}
	}
}
