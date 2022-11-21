package main

import (
	"fmt"
	"math/rand"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func Sixteen() {
	arr := []int{2, 1, -1, 7, 3, 8, 5, 3, -10}

	arr = Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int) []int {
	// если длинна масива = 1, то возвращаем массив
	if len(arr) < 2 {
		return arr
	}
	// левая, правая граница
	left, right := 0, len(arr)-1
	//выбираем случайное число по которому начнется сортировка
	centre := rand.Int() % len(arr)

	arr[centre], arr[right] = arr[right], arr[centre]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	fmt.Println(arr[:left], arr[left+1:])
	Sort(arr[:left])
	Sort(arr[left+1:])

	return arr
}
