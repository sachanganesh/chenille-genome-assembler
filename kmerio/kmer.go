package kmerio

// "github.com/emirpasic/gods/sets/hashset"

// ===================================
// Kmer Utilities
// ===================================

type Kmer struct {
	content	string
}

func NewKmer(kmer string) Kmer {
	return Kmer{content: kmer}
}

func MapNucleotidesToInts(kmer string) []int {
	new_kmer := make([]int, len(kmer))

	for i, nt := range kmer {
		if nt == 'A' {
			new_kmer[i] = 0
		} else if nt == 'C' {
			new_kmer[i] = 1
		} else if nt == 'G' {
			new_kmer[i] = 2
		} else if nt == 'T' {
			new_kmer[i] = 3
		} else {
			new_kmer[i] = 0
		}
	}

	return new_kmer
}

// ===================================
// KmerSet Utilities
// ===================================

type KmerSet struct {
	set map[Kmer]bool
}

func NewKmerSet() KmerSet {
	return KmerSet{set: make(map[Kmer]bool)}
}

func (ks *KmerSet) includes(kmer Kmer) bool {
	return ks.set[kmer]
}

func (ks *KmerSet) add(kmer Kmer) bool {
	if !ks.includes(kmer) {
		ks.set[kmer] = true
		return true
	} else {
		return false
	}
}

func (ks *KmerSet) addAll(kmers []Kmer) []bool {
	added := make([]bool, len(kmers))

	for i := range kmers {
		added[i] = ks.add(kmers[i])
	}

	return added
}

// ===================================
// Kmer Operations
// ===================================

func GetKmersFromShortRead(k int, sr ShortRead) []Kmer {
	num_kmers := len(sr.content) - k + 1
	kmers := make([]Kmer, num_kmers)

	for i := range kmers {
		kmers[i] = NewKmer(sr.content[i : i + k])
	}

	return kmers
}
