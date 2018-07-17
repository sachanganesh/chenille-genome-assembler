package test_time

import (
	"fmt"
	"time"
	"runtime"

	"velour/test"
)

func TimeAll(fragments []string, k int) {
	TimeHMGraph(fragments, k)
	runtime.GC()

	TimeSortedGraph(fragments, k)
	runtime.GC()
}

func TimeHMGraph(fragments []string, k int) {
	start := time.Now()
	test.TestHMGraph(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of HashMap Graph", elapsed)
}

func TimeSortedGraph(fragments []string, k int) {
	start := time.Now()
	test.TestSortedGraph(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of Sorted Graph", elapsed)
}
