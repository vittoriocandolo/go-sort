package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

const (
	sliceLength = 10
	maxValue = 100
)

func main() {
	slice := generateRandomSlice()
	fmt.Println("Input:\t\t\t\t", slice)

	sortedSlice := make([]int, sliceLength)
	copySlice(sortedSlice, slice)

	sortAlgorithms := []struct {
		name string
		fn   func([]int)
	}{
		{"bubble sort", bubbleSort},
		{"selection sort", selectionSort},
		{"insertion sort", insertionSort},
		{"merge sort", func(s []int) { sortedSlice = mergeSort(s) }},
		{"heap sort", heapSort},
		{"quick sort", quickSort},
		{"counting sort", func(s []int) { sortedSlice = countingSort(s, maxValue) }},
		{"tree sort", func(s []int) { sortedSlice = treeSort(s) }},
		{"cycle sort", cycleSort},
		{"bitonic sort", func(s []int) {sortedSlice = bitonicSort(s) }},
		// {"bitonic sort dual", func(s []int) {sortedSlice = bitonicSortDual(s) }},
		{"golang sort", func(s []int) { sort.Slice(s, func(i, j int) bool {
		    return s[i] < s[j]
		})}},
	}
	for _, algorithm := range sortAlgorithms {
	    copySlice(sortedSlice, slice)
	    startTime := time.Now()
		algorithm.fn(sortedSlice)
		elapsedTime := time.Since(startTime)
		fmt.Printf("Sorted with %-*s %v (elapsed time: %e)\n", 20, algorithm.name+":", sortedSlice, elapsedTime.Seconds())
	}
}

func generateRandomSlice() []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, sliceLength)
	for i := range slice {
		slice[i] = rand.Intn(maxValue)
	}
	return slice
}

func copySlice(dst, src []int) {
	for i := range dst {
		dst[i] = src[i]
	}
}
