package kmerio

import (
	"os"
	"bufio"
	"strings"
)

type ShortRead struct {
	Content string
}

func NewShortRead(raw_sr string) ShortRead {
	return ShortRead{Content: strings.ToUpper(raw_sr)}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFastQ(filepath string, k int) KmerSet {
	f, err := os.Open(filepath)
	checkError(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	kmers := NewKmerSet()
	i := 0

	for s.Scan() {
		if (i - 1) % 4 == 0 {
			sr := NewShortRead(s.Text())

			kmers.AddAll(GetKmersFromShortRead(k, sr))
		}
		i++
	}

	return kmers
}
