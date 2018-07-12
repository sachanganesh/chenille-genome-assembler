package test_time

import (
	"fmt"
	"time"
	"runtime"

	"velour/test"
)

func TimeAllTests(fragments []string, k int) {
	TimeSequential(fragments, k)
	runtime.GC()

	TimeSequentialWithLocks(fragments, k)
	runtime.GC()

	TimeConcurrent(fragments, k)
	runtime.GC()
}

func TimeSequential(fragments []string, k int) {
	start := time.Now()
	test.TestSequential(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of Sequential W/O Locks", elapsed)
}
func TimeSequentialWithLocks(fragments []string, k int) {
	start := time.Now()
	test.TestSequentialWithLocks(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of Sequential W/  Locks", elapsed)
}

func TimeConcurrent(fragments []string, k int) {
	start := time.Now()
	test.TestConcurrent(fragments, k)
	elapsed := time.Since(start)

	fmt.Println("Time of Concurrent W/  Locks", elapsed)
}
