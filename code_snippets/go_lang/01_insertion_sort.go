package main

import (
	"cmp"
	"fmt"
)

func main() {

	fmt.Println("--- insertion sort wiht integers")
	arr := []int{5, 32, 7, 2, 86, 54, 100}

	fmt.Println("Before: ", arr)
	insertionSort(arr)
	fmt.Println("After: ", arr)

	fmt.Println("--- insertion sort with Words")
	arrSlice := []string{"gold", "apple", "orange", "money"}
	fmt.Println("Before: ", arrSlice)
	insertionSort(arrSlice)
	fmt.Println("After: ", arrSlice)
}

func insertionSort[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}
