package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func Eleven() {
	arrOne := []int{1, 2, 33, 44, 7, 8, 9, 0}
	arrTwo := []int{4, 2, 5, 7, 8, 22, 55, 77, 98, 1}
	ArrRes := make([]int, 0)
	for i := 0; i < len(arrOne); i++ {
		for j := 0; j < len(arrTwo); j++ {
			if arrOne[i] == arrTwo[j] {
				ArrRes = append(ArrRes, arrOne[i])
			}
		}
	}
	fmt.Println(ArrRes)

	Map := make(map[int]int)
	for _, value := range arrOne {
		Map[value] = 1
	}
	for _, value := range arrTwo {
		if _, ok := Map[value]; ok {
			Map[value] = Map[value] + 1
		} else {
			Map[value] = 1
		}
	}
	resMap := make([]int, 0)
	for key, value := range Map {
		if value > 1 {
			resMap = append(resMap, key)
		}
	}
	fmt.Println(resMap, Map)
}
