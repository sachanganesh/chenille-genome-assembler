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

	TimePASHAGraph(fragments, k)
	runtime.GC()
}

func TimeHMGraph(fragments []string, k int) {
	start := time.Now()
	test.TestHMGraph(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of HashMap Graph", elapsed)
}

func TimePASHAGraph(fragments []string, k int) {
	start := time.Now()
	test.TestPASHAGraph(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of PASHA Graph", elapsed)
}
