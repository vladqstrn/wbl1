package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
func Thirteen() {
	a := 8123
	b := 2147483647

	FirstSwap(a, b)
	SecondSwap(a, b)
}

func FirstSwap(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

func SecondSwap(a, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
