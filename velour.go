package main

import (
	"fmt"
	"velour/kmerio"
	"time"
)

func main() {
	filepath := "data/staphylococcus_aureus/frag1.fastq"
	kmer_len := 7

	start := time.Now()
	kmerio.ParseFastQ(filepath, kmer_len)
	elapsed := time.Since(start)

	fmt.Println("Time elapsed %s", elapsed)
}
