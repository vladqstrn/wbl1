package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(22+32+42....) с использованием конкурентных вычислений.
func Three() {
	var arr = []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	sum := 0
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		value := arr[i]
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			sum = sum + (value * value)

		}()
	}
	wg.Wait()
	fmt.Println(sum)

}
