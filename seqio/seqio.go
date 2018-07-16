package seqio

import (
	"os"
	"bufio"
	"strings"
	"math"

	"velour/debruijn"
)

func checkError(e error) {
	if e != nil {
		panic(e)
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

// ===================================
// ShortRead
// ===================================

type ShortRead string

func NewShortRead(raw_sr string) ShortRead {
	return ShortRead(strings.ToUpper(raw_sr))
}

func (sr ShortRead) Len() int {
	return len(sr)
}

func (sr ShortRead) Substring(start int, end_before int) string {
	return string(sr[start : end_before])
}

func (sr ShortRead) GetKmers(k int) []debruijn.Kmer {
	num_kmers := sr.Len() - k + 1
	kmers := make([]debruijn.Kmer, num_kmers)

	for i := range kmers {
		kmers[i] = debruijn.NewKmer(sr.Substring(i, i + k))
	}

	return kmers
}

// ===================================
// IO Operations
// ===================================

func GraphFromFastQ(fragment string, k int, graph debruijn.Graph) debruijn.Graph {
	f, err := os.Open(fragment)
	checkError(err)
	defer f.Close()

	s := bufio.NewScanner(f)

	i := 0
	tmp := 0
	for s.Scan() {
		if (i - 1) % 4 == 0 {
			tmp++
			sr := NewShortRead(s.Text())
			graph.AddNodes(sr.GetKmers(k))

			if tmp == 10000 {
				return graph
			}
		}

		i++
	}

	return graph
}
