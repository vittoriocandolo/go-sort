package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println("Input:", slice)

	// Sort slice with bubble sort
	bubbleSorted := make([]int, len(slice))
	copy(bubbleSorted, slice)
	bubbleSort(bubbleSorted)
	fmt.Println("Sorted with bubble sort:\t", bubbleSorted)

	// Sort slice with selection sort
	selectionSorted := make([]int, len(slice))
	copy(selectionSorted, slice)
	selectionSort(selectionSorted)
	fmt.Println("Sorted with selection sort:\t", selectionSorted)

	// Sort slice with insertion sort
	insertionSorted := make([]int, len(slice))
	copy(insertionSorted, slice)
	insertionSort(insertionSorted)
	fmt.Println("Sorted with insertion sort:\t", insertionSorted)

	// Sort slice with merge sort
	mergeSorted := make([]int, len(slice))
	mergeSorted = mergeSort(slice)
	fmt.Println("Sorted with merge sort:\t\t", mergeSorted)

	// Sort slice with heap sort
	heapSorted := make([]int, len(slice))
	copy(heapSorted, slice)
	heapSort(heapSorted)
	fmt.Println("Sorted with heap sort:\t\t", heapSorted)

	// Sort slice with quick sort
	quickSorted := make([]int, len(slice))
	copy(quickSorted, slice)
	quickSort(quickSorted)
	fmt.Println("Sorted with quick sort:\t\t", quickSorted)

	// Sort slice with counting sort
	countingSorted := countingSort(slice, 100)
	fmt.Println("Sorted with counting sort:\t", countingSorted)

	// Sort slice with tree sort
	treeSorted := make([]int, len(slice))
	treeSorted = treeSort(slice)
	fmt.Println("Sorted with tree sort:\t\t", treeSorted)

	// Sort slice with cycle sort
	cycleSorted := make([]int, len(slice))
	copy(cycleSorted, slice)
	cycleSort(cycleSorted)
	fmt.Println("Sorted with cycle sort:\t\t", cycleSorted)

	// Sort slice with bitonic sort
	
}
