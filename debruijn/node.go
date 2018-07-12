package debruijn

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

// ===================================
// KmerNode
// ===================================

type KmerNode struct {
	kmer			string
	frequency		int
	predecessors	[]int
	successors		[]int
}

func NewNode(kmer string) *KmerNode {
	return &KmerNode{kmer, 1, make([]int, 0, 2), make([]int, 0, 2)}
}

// ===================================
// KmerNode Functions
// ===================================

func (kmer_entry *KmerNode) GetKmer() string {
	return kmer_entry.kmer
}

func (kmer_entry *KmerNode) GetFrequency() int {
	return kmer_entry.frequency
}

func (kmer_entry *KmerNode) SetFrequency(new_frequency int) {
	kmer_entry.frequency = new_frequency
}

func (kmer_entry *KmerNode) IncrementFrequency() {
	kmer_entry.SetFrequency(kmer_entry.GetFrequency() + 1)
}

func (kmer_entry *KmerNode) GetPredecessors() []int {
	return kmer_entry.predecessors
}

func (kmer_entry *KmerNode) IsAPredecessor(needle int) bool {
	for _, pred := range kmer_entry.predecessors {
		if pred == needle {
			return true
		}
	}

	return false
}

func (kmer_entry *KmerNode) AddPredecessor(pred_id int) {
	if !kmer_entry.IsAPredecessor(pred_id) {
		kmer_entry.predecessors = append(kmer_entry.predecessors, pred_id)
	}
}

func (kmer_entry *KmerNode) GetSuccessors() []int {
	return kmer_entry.successors
}

func (kmer_entry *KmerNode) IsASuccessor(needle int) bool {
	for _, succ := range kmer_entry.successors {
		if succ == needle {
			return true
		}
	}

	return false
}

func (kmer_entry *KmerNode) AddSuccessor(succ_id int) {
	if !kmer_entry.IsASuccessor(succ_id) {
		kmer_entry.successors = append(kmer_entry.successors, succ_id)
	}
}

func (kmer_entry *KmerNode) Merge(other_entry GraphNode) {
	if kmer_entry.GetKmer() == other_entry.GetKmer() {	kmer_entry.SetFrequency(kmer_entry.GetFrequency() + other_entry.GetFrequency())

		for _, pred_id := range other_entry.GetPredecessors() {
			kmer_entry.AddPredecessor(pred_id)
		}

		for _, succ_id := range other_entry.GetSuccessors() {
			kmer_entry.AddSuccessor(succ_id)
		}
	}
}
