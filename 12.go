package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
//собственное множество.
func Twelve() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	Map := make(map[string]bool)

	for _, i := range arr {
		Map[i] = true
	}

	keys := make([]string, 0, len(Map))
	for k := range Map {
		keys = append(keys, k)
	}

	fmt.Println(keys)

}
