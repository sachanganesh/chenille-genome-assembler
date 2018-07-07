package kmerio

import (
	// "github.com/emirpasic/gods/sets/hashset"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getKmersFromString(k int, seq []byte) [][]byte {
	num_kmers := len(seq) - k + 1
	kmers := make([][]byte, num_kmers)

	for i := range kmers {
		kmers[i] = make([]byte, k)
		copy(kmers[i], seq[i : i + k])
	}

	return kmers
}
