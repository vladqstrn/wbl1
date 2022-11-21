package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	c  int
	mu sync.Mutex
	wg sync.WaitGroup
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c += 1
	c.wg.Done()
}

func CreateCounter() *Counter {
	c := 0
	return &Counter{
		c,
		sync.Mutex{},
		sync.WaitGroup{},
	}
}

func StartInc(counter *Counter) {
	counter.Increment()
}

func Eighteen() {
	counter := CreateCounter()

	counter.wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go StartInc(counter)
	}
	counter.wg.Wait()

	fmt.Println(counter.c)

	time.Sleep(4 * time.Second)
}
