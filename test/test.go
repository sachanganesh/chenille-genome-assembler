package test

import (
	"fmt"
	"sync"
	"unsafe"

	"velour/kmerio"
	"velour/debruijn"
	"velour/debruijn/sync"
)

func printGraphData(graph debruijn.Graph) {
	fmt.Println("Number of Unique Kmers:", graph.Len())
	fmt.Println("Size of graph:", unsafe.Sizeof(graph))
}

func TestSequential(fragments []string, k int) debruijn.Graph {
	fmt.Println("\nTesting Sequential Graph Read")

	var graph debruijn.Graph = debruijn.NewGraph()

	for _, fragment := range fragments {
		kmerio.GraphFromFastQ(fragment, k, graph)
	}

	printGraphData(graph)
	return graph
}

func TestSequentialWithLocks(fragments []string, k int) debruijn.Graph {
	fmt.Println("\nTesting Sequential with Locks Graph Read")

	var graph debruijn.Graph = debruijn_sync.NewGraph()

	for _, fragment := range fragments {
		kmerio.GraphFromFastQ(fragment, k, graph)
	}

	printGraphData(graph)
	return graph
}

func TestConcurrent(fragments []string, k int) debruijn.Graph {
	fmt.Println("\nTesting Concurrent Graph Read")

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

	printGraphData(graph)
	return graph
}
