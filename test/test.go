package test

import (
	"fmt"
	// "sync"

	"velour/seqio"
	"velour/debruijn"
	"velour/debruijn/hmgraph"
	"velour/debruijn/pasha"
)

func printGraphData(graph debruijn.Graph) {
	fmt.Println("Number of Total  Kmers:", graph.GetNumNodesSeen())
	fmt.Println("Number of Unique Kmers:", graph.Len())
}

func TestHMGraph(fragments []string, k int) debruijn.Graph {
	fmt.Println("\nTesting HashMap Graph Read")

	var node_gen debruijn.NodeGenerator = hmgraph.NewNode
	var graph debruijn.Graph = hmgraph.NewGraph(node_gen)

	for _, fragment := range fragments {
		seqio.GraphFromFastQ(fragment, k, graph)
	}

	printGraphData(graph)
	return graph
}

func TestPASHAGraph(fragments []string, k int) debruijn.Graph {
	fmt.Println("\nTesting PASHA Graph Read")

	var node_gen debruijn.NodeGenerator = pasha.NewNode
	var graph debruijn.Graph = pasha.NewGraph(node_gen)

	for _, fragment := range fragments {
		seqio.GraphFromFastQ(fragment, k, graph)
	}

	printGraphData(graph)
	return graph
}

// func TestSequentialWithLocks(fragments []string, k int) debruijn.Graph {
// 	fmt.Println("\nTesting Sequential with Locks Graph Read")
//
// 	var graph debruijn.Graph = debruijn.debruijn_sync.NewGraph(debruijn.hmgraph.NewGraph())
//
// 	for _, fragment := range fragments {
// 		kmerio.GraphFromFastQ(fragment, k, graph)
// 	}
//
// 	printGraphData(graph)
// 	return graph
// }
//
// func TestConcurrent(fragments []string, k int) debruijn.Graph {
// 	fmt.Println("\nTesting Concurrent Graph Read")
//
// 	var graph debruijn.Graph = debruijn_sync.NewGraph()
// 	wg := &sync.WaitGroup{}
//
// 	for _, fragment := range fragments {
// 		wg.Add(1)
//
// 		go func(fragment string) {
// 			defer wg.Done()
//
// 			kmerio.GraphFromFastQ(fragment, k, graph)
// 		}(fragment)
// 	}
//
// 	wg.Wait()
//
// 	printGraphData(graph)
// 	return graph
// }
