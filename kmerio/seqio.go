package kmerio

import (
	"os"
	"bufio"
	"strings"

	"velour/debruijn"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// ===================================
// ShortRead Utilities
// ===================================

type ShortRead string

func NewShortRead(raw_sr string) ShortRead {
	return strings.ToUpper(raw_sr)}
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
	for s.Scan() {
		if (i - 1) % 4 == 0 {
			sr := NewShortRead(s.Text())
			graph.AddNodes(GetKmersFromShortRead(k, sr))
		}

		i++
	}

	return graph
}
