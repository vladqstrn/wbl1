package main

import (
	"fmt"
	"math"
	"strings"
)

func Ten() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := make(map[int][]float64)
	for i := 0; i < len(temp); i++ {
		buff := math.Floor(temp[i]/10) * 10
		tempMap[int(buff)] = append(tempMap[int(buff)], temp[i])
	}
	var keys []int
	for k := range tempMap {
		keys = append(keys, k)
	}

	var str string
	for _, b := range keys {
		s := fmt.Sprintf("%f", tempMap[b])
		str += fmt.Sprint(b) + ":" + s + " "
	}
	badStr := []string{"[", "]"}
	goodStr := []string{"{", "}"}
	for i := 0; i < 2; i++ {
		str = strings.Replace(str, badStr[i], goodStr[i], -1)
	}

	fmt.Println(str)

}
