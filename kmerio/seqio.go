package kmerio

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

type ShortRead struct {
	content string
}

func NewShortRead(raw_sr string) ShortRead {
	return ShortRead{content: strings.ToUpper(raw_sr)}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFastQ(filepath string, k int) {
	f, err := os.Open(filepath)
	checkError(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	kmers := NewKmerSet()
	i := 0

	for s.Scan() {
		if (i - 1) % 4 == 0 {
			sr := NewShortRead(s.Text())

			kmers.addAll(GetKmersFromShortRead(k, sr))
		}
		i++
	}

	fmt.Println(len(kmers.set))
}
