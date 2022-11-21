package main

import "fmt"

//Удалить i-ый элемент из слайса.

func TwentyThree() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	valueDel := 3

	arr[valueDel] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}
