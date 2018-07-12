package main

import (
	"fmt"
	"time"
	"sync"

	"velour/kmerio"
	"velour/debruijn"
	"velour/debruijn/sync"
)

func main() {
	fragments := []string{"data/staphylococcus_aureus/frag1.fastq", "data/staphylococcus_aureus/frag2.fastq"}

	k := kmerio.EstimateK(2900000)
	fmt.Println("Estimated K:", k)

	timeConcurrentTest(fragments, k)
	timeAllTests(fragments, k)
}

func timeAllTests(fragments []string, k int) {
	start := time.Now()

	testSequential(fragments, k)
	t1 := time.Now()

	testSequentialWithLocks(fragments, k)
	t2 := time.Now()

	testConcurrent(fragments, k)
	t3 := time.Now()

	fmt.Println("Time for Sequential W/O Locks", t1.Sub(start))
	fmt.Println("Time for Sequential W/  Locks", t2.Sub(t1))
	fmt.Println("Time for Concurrent W/  Locks", t3.Sub(t2))
}

func timeConcurrentTest(fragments []string, k int) {
	start := time.Now()
	graph := testConcurrent(fragments, k)
	elapsed := time.Since(start)

	fmt.Println(graph.Len(), "Number of Unique Kmers")
	fmt.Println("Time elapsed", elapsed)
}

func testSequential(fragments []string, k int) debruijn.Graph {
	var graph debruijn.Graph = debruijn.NewGraph()

	for _, fragment := range fragments {
		kmerio.GraphFromFastQ(fragment, k, graph)
	}

	return graph
}

func testSequentialWithLocks(fragments []string, k int) debruijn.Graph {
	var graph debruijn.Graph = debruijn_sync.NewGraph()

	for _, fragment := range fragments {
		kmerio.GraphFromFastQ(fragment, k, graph)
	}

	return graph
}

func testConcurrent(fragments []string, k int) debruijn.Graph {
	var graph debruijn.Graph = debruijn_sync.NewGraph()
	wg := &sync.WaitGroup{}

	for _, fragment := range fragments {
		wg.Add(1)

		go func(fragment string) {
			defer wg.Done()

			kmerio.GraphFromFastQ(fragment, k, graph)
		}(fragment)
	}

	wg.Wait()
	return graph
}
