package main

import (
	"fmt"
	"time"

	"velour/kmerio"
	"velour/debruijn"
)

func main() {
	fragments := []string{"data/staphylococcus_aureus/frag1.fastq", "data/staphylococcus_aureus/frag2.fastq"}

	k := kmerio.EstimateK(2900000)
	fmt.Println("Estimated K:", k)

	start := time.Now()

	var graph debruijn.Graph = debruijn.NewGraph()

	for _, fragment := range fragments {
		graph = kmerio.GraphFromFastQ(fragment, k, graph)
	}

	elapsed := time.Since(start)

	fmt.Println("Number of Kmers Tracked:", graph.Len())
	fmt.Println("Total time elapsed", elapsed)
}
