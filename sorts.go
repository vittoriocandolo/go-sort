package main

import (
    "math"
    "sync"
)


func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if slice[j] > slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}
}


func selectionSort(slice []int) {
	n := len(slice)
	for i := 0; i < n - 1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
}


func insertionSort(slice []int) {
	n := len(slice)
	for j := 1; j < n; j++ {
		k := slice[j]
		i := j - 1
		for i >= 0 && slice[i] > k {
			slice[i + 1] = slice[i]
			i--
		}
		slice[i + 1] = k
	}
}


func mergeSort(slice []int) []int {
	n := len(slice)
	if n <= 1 {
		return slice
	}
	mid := n / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			result = append(result, right...)
			break
		}
		if len(right) == 0 {
			result = append(result, left...)
			break
		}
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}


func heapSort(slice []int) {
	n := len(slice)
	for i := n / 2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, i, 0)
	}
}

func heapify(slice []int, n, i int) {
	largest := i
	left := 2 * i + 1
	right := 2 * i + 2
	if left < n && slice[left] > slice[largest] {
		largest = left
	}
	if right < n && slice[right] > slice[largest] {
		largest = right
	}
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapify(slice, n, largest)
	}
}


func quickSort(slice []int) {
	quickSortRecursive(slice, 0, len(slice) - 1)
}

func quickSortRecursive(slice []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition(slice, left, right)
	quickSortRecursive(slice, left, pivotIndex - 1)
	quickSortRecursive(slice, pivotIndex + 1, right)
}

func partition(slice []int, left, right int) int {
	pivotIndex := right
	pivotValue := slice[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if slice[i] < pivotValue {
			slice[storeIndex], slice[i] = slice[i], slice[storeIndex]
			storeIndex++
		}
	}
	slice[storeIndex], slice[pivotIndex] = slice[pivotIndex], slice[storeIndex]
	return storeIndex
}


func countingSort(slice []int, maxValue int) []int {
	n := len(slice)
	count := make([]int, maxValue + 1)
	for i := 0; i < n; i++ {
		count[slice[i]]++
	}
	sorted := make([]int, n)
	sortedIndex := 0
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			sorted[sortedIndex] = i
			sortedIndex++
		}
	}

	return sorted
}


type TreeNode struct {
	key int
	left  *TreeNode
	right *TreeNode
}

func treeSort(slice []int) []int {
	var root *TreeNode
	for _, v := range slice {
		root = insert(root, v)
	}
	sorted := make([]int, 0, len(slice))
	traverse(root, &sorted)
	return sorted
}

func insert(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return &TreeNode{key: key}
	}
	if key < node.key {
		node.left = insert(node.left, key)
	} else {
		node.right = insert(node.right, key)
	}
	return node
}

func traverse(node *TreeNode, slice *[]int) {
	if node == nil {
		return
	}
	traverse(node.left, slice)
	*slice = append(*slice, node.key)
	traverse(node.right, slice)
}


func cycleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n - 1; i++ {
		item := slice[i]
		pos := i
		for j := i + 1; j < n; j++ {
			if slice[j] < item {
				pos++
			}
		}
		if pos == i {
			continue
		}
		for item == slice[pos] {
			pos++
		}
		if pos != i {
			item, slice[pos] = slice[pos], item
		}
		for pos != i {
			pos = i
			for j := i + 1; j < n; j++ {
				if slice[j] < item {
					pos++
				}
			}
			for item == slice[pos] {
				pos++
			}
			if item != slice[pos] {
				item, slice[pos] = slice[pos], item
			}
		}
	}
}

func bitonicSort(slice []int) []int {
    n := len(slice)
    makePowerOfTwo(&slice)
	bitonicSortRec(slice, 0, len(slice), 1)
	return slice[:n]
}

func makePowerOfTwo(slice *[]int) {
	n := len(*slice)
    if (isPowerOfTwo(n)) {
        return
    } else {
        for len(*slice) < nextPowerOfTwo(n) {
            *slice = append(*slice, math.MaxInt32)
        }
    }
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func nextPowerOfTwo(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}

func bitonicSortRec(slice []int, low, cnt, dir int) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSortRec(slice, low, k, 1)
		bitonicSortRec(slice, low+k, k, 0)
		bitonicMerge(slice, low, cnt, dir)
	}
}

func bitonicMerge(slice []int, low, cnt, dir int) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			compAndSwap(slice, i, i+k, dir)
		}
		bitonicMerge(slice, low, k, dir)
		bitonicMerge(slice, low+k, k, dir)
	}
}

func compAndSwap(slice []int, i, j, dir int) {
	if ((dir==1 && slice[i] > slice[j]) || (dir==0 && slice[i] < slice[j])) {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// only if len(slice) is a power of two...
func bitonicSortDual(slice []int) []int {
    n := len(slice)
    makePowerOfTwo(&slice)

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        bitonicSortRec(slice, 0, len(slice)/2, 1)
    }()
    go func() {
        defer wg.Done()
        bitonicSortRec(slice, len(slice)/2, len(slice)/2, 0)
    }()

    wg.Wait()
    return slice[:n]
}
