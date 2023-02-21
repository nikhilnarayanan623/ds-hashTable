package main

import "fmt"

func main() {

	// m := hashMap.GetHashMap()
	// m.Put("Nikhil", "Hello")
	// m.Display()

	// array := []int{1, 35, 2, 6, 4, 76}

	// test(1)

	// fmt.Println(array)
}

func quickSort(array []int) {
	quickSortHelper(array, 0, len(array)-1)
}

func quickSortHelper(array []int, start int, end int) {

	if start >= end {
		return
	}

	pIdx := partition(array, start, end)

	quickSortHelper(array, start, pIdx)
	quickSortHelper(array, pIdx+1, end)
}

func partition(array []int, start int, end int) int {

	pIdx := start - 1

	for i := start; i < end; i++ {

		if array[i] < array[end] {
			pIdx++

			array[pIdx], array[i] = array[i], array[pIdx]
		}
	}

	pIdx++
	array[pIdx], array[end] = array[end], array[pIdx]

	return pIdx
}

func mergeSort(array []int) {

	mergeSortHelper(array, 0, len(array)-1)
}

func mergeSortHelper(array []int, start int, end int) {

	if start >= end {
		return
	}

	mid := start + ((end - start) / 2)

	mergeSortHelper(array, start, mid)
	mergeSortHelper(array, mid+1, end)

	merge(array, start, mid, end)
}

func merge(array []int, start int, mid int, end int) {

	newArray := make([]int, (end-start)+1)
	newIdx := 0

	leftIdx := start
	rightIdx := mid + 1

	// copy the merged array two sorted portion to new array
	for leftIdx <= mid || rightIdx <= end {

		if leftIdx <= mid && rightIdx <= end {

			if array[leftIdx] < array[rightIdx] {
				newArray[newIdx] = array[leftIdx]
				leftIdx++
			} else {
				newArray[newIdx] = array[rightIdx]
				rightIdx++
			}
		} else if leftIdx <= mid {
			newArray[newIdx] = array[leftIdx]
			leftIdx++
		} else {
			newArray[newIdx] = array[rightIdx]
			rightIdx++
		}
		newIdx++

	}
	//copy the new arrays value to the old array position

	leftIdx = start

	for _, val := range newArray {

		array[leftIdx] = val
		leftIdx++
	}

}

func test(n int) {

	var a, b = 1, 2
	a = b
	b = a

	fmt.Println(a, b, "\n", n)
	test(n + 1)
}
