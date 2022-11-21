package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.
func Seventeen() {
	arr := []int{0, 1, 4, 6, 7, 8, 9}
	fmt.Println(BinarySearch(arr, 8))
}

func BinarySearch(arr []int, key int) (bool, int) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		centre := (low + high) / 2
		fmt.Println(centre)

		if arr[centre] < key {
			low = centre + 1
		} else {
			high = centre - 1
		}
	}
	if low == len(arr) || arr[low] != key {
		return false, 0
	}

	return true, low
}
