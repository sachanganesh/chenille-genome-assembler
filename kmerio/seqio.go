package kmerio

import (
	"os"
	"bufio"
	"strings"

	"velour/debruijn"
)

type ShortRead struct {
	Sequence string
}

func NewShortRead(raw_sr string) ShortRead {
	return ShortRead{Sequence: strings.ToUpper(raw_sr)}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func GraphFromFastQ(filepath string, k int) debruijn.DBGraph {
	f, err := os.Open(filepath)
	checkError(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	graph := debruijn.NewDBGraph()

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
