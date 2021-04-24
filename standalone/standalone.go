package main

import (
	"fmt"
	"log"
	"os"
	_runtime "runtime/pprof"
)

var (
	cpuProfile = "cpu.prof"
)

func main() {
	// Create file to stored the profiling
	cpuProf, err := os.Create(cpuProfile)
	if err != nil {
		log.Fatal(err)
	}
	defer cpuProf.Close()
	_runtime.StartCPUProfile(cpuProf)
	defer _runtime.StopCPUProfile()
	ProcessBigArray()

}

func ProcessBigArray() {
	for i := 0; i < 10; i++ {
		arr := bigArray()
		if arr == nil {
			fmt.Println("Array is Nil")
		}
	}
}

func bigArray() *[]int {
	s := make([]int, 10)
	return &s
}
