package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func QuickSortGuard(arr []int) ([]int, error) {
	if len(arr) <= 1 {
		return arr, errors.New("Short array")
	} else {
		sortedArray := QuickSort(arr)
		return sortedArray, nil
	}
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func main() {
	arr := []int{9, 2, 10, 12, 3, 7, 8, 4, 6, 5, 1}
	fmt.Println("Initial array is:", arr)

	sortedArr, err := QuickSortGuard(arr)
	if err != nil {
		fmt.Println("Error occured: " + err.Error())
	} else {
		fmt.Println("Sorted array is:", sortedArr)
	}
}
