package main

import (
	"fmt"
	"sync"
	"time"

	"velour/kmerio"
)

func main() {
	fragments := [...]string{"data/staphylococcus_aureus/frag1.fastq", "data/staphylococcus_aureus/frag2.fastq"}

	k := kmerio.EstimateK(2900000)
	fmt.Println("Estimated K:", k)

	start := time.Now()

	kmers_chan := make(chan kmerio.KmerSet)
	defer close(kmers_chan)

	var frag_wg sync.WaitGroup
	frag_wg.Add(len(fragments))

	// var kmers kmerio.KmerSet
	kmers := kmerio.NewKmerSet()
	fmt.Println("Starting with:", len(kmers.Set), "kmers")

	// frag1 := kmerio.ParseFastQ(fragments[0], k)
	// fmt.Println("Fragment 1 Size:", len(frag1.Set))
	// kmers.Merge(frag1)
	//
	// frag2 := kmerio.ParseFastQ(fragments[1], k)
	// fmt.Println("Fragment 2 Size:", len(frag2.Set))
	// kmers.Merge(frag2)

	for i := 0; i < len(fragments); i++ {
		go func(i int) {
			defer frag_wg.Done()
			kmers_chan <- kmerio.ParseFastQ(fragments[i], k)
		}(i)
	}

	for ks := range kmers_chan {
		kmers.Merge(ks)
	}

	frag_wg.Wait()

	elapsed := time.Since(start)

	fmt.Println("Ending with:", len(kmers.Set), "kmers")
	fmt.Println("Time elapsed", elapsed)

	// for kmer := range kmers.set {
	// 	fmt.Println(kmer, kmers.set[kmer])
	// }
}
