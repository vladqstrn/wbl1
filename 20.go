package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func Twenty() {
	str := "snow dog sun"
	words := strings.Fields(str)
	arr := make([]string, 0)
	for i := len(words) - 1; i >= 0; i-- {
		arr = append(arr, words[i])
	}
	fmt.Println(arr)
	sentence := strings.Join(arr, " ")
	fmt.Println(sentence)
}
