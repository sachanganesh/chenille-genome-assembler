package main

import (
	"fmt"

	"velour/kmerio"
	"velour/test/time"
)

func main() {
	fragments := []string{"data/staphylococcus_aureus/frag1.fastq", "data/staphylococcus_aureus/frag2.fastq"}

	k := kmerio.EstimateK(2900000)
	fmt.Println("Estimated K:", k)

	test_time.TimeAll(fragments, k)
}
