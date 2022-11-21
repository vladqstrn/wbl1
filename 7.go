package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

type goMap struct {
	mutex sync.Mutex
	m     map[int]int
}

func GetMap() *goMap {
	myMap := make(map[int]int)
	return &goMap{
		sync.Mutex{},
		myMap,
	}
}

// геттер
func (g *goMap) Set(key, value int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.m[key] = value

}

// сеттер
func (g *goMap) Get(key int) (int, bool) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	val, ok := g.m[key]
	return val, ok
}

func pushToMap(value int, myMap *goMap) {
	myMap.Set(value, value)
}

func Seven() {
	myMap := GetMap()

	for i := 0; i < 10; i++ {
		go pushToMap(i, myMap)
	}

	time.Sleep(1 * time.Second)
	for i := range myMap.m {
		fmt.Println(myMap.Get(i))
	}
}
