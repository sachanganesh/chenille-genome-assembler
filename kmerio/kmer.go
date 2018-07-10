package kmerio

// "github.com/emirpasic/gods/sets/hashset"

import (
	"math"
)

// ===================================
// Kmer Utilities
// ===================================

func CleanKmer(kmer string) string {
	for i, nt := range kmer {
		if nt != 'A' && nt != 'C' && nt != 'G' && nt != 'T' {
			tmp := []byte(kmer)
			tmp[i] = 'A'
			kmer = string(tmp)
		}
	}

	return kmer
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

func GetKmersFromShortRead(k int, sr ShortRead) []string {
	num_kmers := len(sr.Sequence) - k + 1
	kmers := make([]string, num_kmers)

	for i := range kmers {
		kmers[i] = sr.Sequence[i : i + k]
	}

	return kmers
}
