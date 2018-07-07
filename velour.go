package main

import (
	// "fmt"
	"velour/kmerio"
)

func main() {
	filepath := "data/staphylococcus_aureus/frag1.fastq"
	kmer_len := 7

	kmerio.ParseFastQ(filepath, kmer_len)
}
