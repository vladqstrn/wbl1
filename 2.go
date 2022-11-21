package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func Second() {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	var arr = [5]int{2, 4, 6, 8, 10}
	resArr := make([]int, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1) //Счетчик. Если счетчик становится равным нулю, все горутины, заблокированные в ожидании, освобождаются.
		go calcSquare(arr[i], i, resArr, &m, &wg)

	}
	wg.Wait() // ждем пока счетчик не будет равен 0
	fmt.Println(resArr)
	secondArr := Second2()
	fmt.Println(secondArr)
}

func calcSquare(value int, i int, resArr []int, m *sync.Mutex, wg *sync.WaitGroup) {
	m.Lock()
	defer m.Unlock()
	resArr[i] = value * value
	wg.Done() //Done уменьшает значение счетчика WaitGroup на единицу.
}

//======================================================================

func squares(c chan int, arr [5]int) {
	for i := 0; i < len(arr); i++ {
		c <- arr[i] * arr[i]
	}

	close(c) // close channel
}

func Second2() []int {
	c := make(chan int)
	var arr = [5]int{2, 4, 6, 8, 10}

	for i := 0; i < len(arr); i++ {

		go squares(c, arr)
	}
	tempArr := make([]int, 0)

	for val := range c {
		tempArr = append(tempArr, val)
	}
	return tempArr
}
