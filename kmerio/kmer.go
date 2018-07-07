package kmerio

import (
	"os"
	"bufio"
	"strings"
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

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
	var kmers [][]int
	i := 0

	for s.Scan() {
		if (i - 1) % 4 == 0 {
			sr := strings.ToUpper(s.Text())

			for _, kmer := range GetKmersFromShortRead(k, []byte(sr)) {
				kmers = append(kmers, kmer)
			}
		}
		i++
	}

	fmt.Println(len(kmers))
}

func GetKmersFromShortRead(k int, seq []byte) [][]int {
	num_kmers := len(seq) - k + 1
	kmers := make([][]int, num_kmers)

	for i := range kmers {
		kmers[i] = NucleotideToInt(seq[i : i + k])
	}

	return kmers
}

func NucleotideToInt(kmer []byte) []int {
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
