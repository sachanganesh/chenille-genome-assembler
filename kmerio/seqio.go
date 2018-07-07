package kmerio

import (
	"os"
	"bufio"
	"fmt"
)

func ParseFastQ(filepath string, k int) {
	f, err := os.Open(filepath)
	checkError(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	i := 0
	var kmers [][]byte

	for s.Scan() {
		if (i - 1) % 4 == 0 {
			for _, kmer := range getKmersFromString(k, s.Bytes()) {
				kmers = append(kmers, kmer)
			}

			fmt.Println(len(kmers))
		}
		i++
	}
}
