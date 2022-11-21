package main

import (
	"fmt"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func Nine() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	arr := make([]int, 10000)
	//arr = append(arr, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 119, 20)
	for i := range arr {
		arr[i] = i
	}
	counter := 0
	go func() {
		for i := 0; i < len(arr); i++ {
			firstChan <- arr[i]
			counter++
		}
	}()

	go func() {
		defer close(secondChan)
		for elem := range firstChan {
			secondChan <- elem * 2
		}
	}()

	go func() {
		defer close(firstChan)
		for elemm := range secondChan {
			fmt.Println(elemm)
			counter--
		}
	}()
	time.Sleep(10 * time.Second)
	println(counter, " - counter")
}
