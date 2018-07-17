package main

import (
	"fmt"

	"velour/seqio"
	"velour/test/time"

	// "github.com/pkg/profile"
)

func main() {
	// defer profile.Start().Stop()

	fragments := []string{"data/staphylococcus_aureus/frag1.fastq", "data/staphylococcus_aureus/frag2.fastq"}

	k := seqio.EstimateK(2900000)
	fmt.Println("Estimated K:", k)

	test_time.TimeAll(fragments[:1], k)
}
