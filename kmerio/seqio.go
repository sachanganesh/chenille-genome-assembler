package kmerio

import (
	"os"
	"bufio"
	"strings"
	// "sync"

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

type ShortRead struct {
	Sequence string
}

func NewShortRead(raw_sr string) ShortRead {
	return ShortRead{Sequence: strings.ToUpper(raw_sr)}
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
