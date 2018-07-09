package kmerio

// "github.com/emirpasic/gods/sets/hashset"

import (
	"math"
)

// ===================================
// Kmer Utilities
// ===================================

type Kmer struct {
	Content	string
}

func NewKmer(kmer string) Kmer {
	return Kmer{Content: kmer}
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
	Set map[Kmer]int
}

func NewKmerSet() KmerSet {
	return KmerSet{Set: make(map[Kmer]int)}
}

func (ks *KmerSet) Includes(kmer Kmer) bool {
	return ks.Set[kmer] > 0
}

func (ks *KmerSet) Add(kmer Kmer) bool {
	if !ks.Includes(kmer) {
		ks.Set[kmer]++
		return true
	} else {
		ks.Set[kmer]++
		return false
	}
}

func (ks *KmerSet) AddAll(kmers []Kmer) []bool {
	added := make([]bool, len(kmers))

	for i := range kmers {
		added[i] = ks.Add(kmers[i])
	}

	return added
}

func (ks *KmerSet) Merge(os KmerSet) {
	if len(ks.Set) > 0 {
		for other_key := range os.Set {
			ks.Set[other_key] += os.Set[other_key]
		}
	} else {
		ks.Set = os.Set
	}
}

// ===================================
// Kmer Operations
// ===================================

func EstimateK(genome_size int) int {
	k := 1.0
	for {
		if math.Pow(4.0, k) >= float64(genome_size * 2 * 10) {
			return int(k)
		}

		k++
	}
}

func GetKmersFromShortRead(k int, sr ShortRead) []Kmer {
	num_kmers := len(sr.Content) - k + 1
	kmers := make([]Kmer, num_kmers)

	for i := range kmers {
		kmers[i] = NewKmer(sr.Content[i : i + k])
	}

	return kmers
}
