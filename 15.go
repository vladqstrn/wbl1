package main

import "fmt"

// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
// func main() {
// 	someFunc()
// }

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string

func Fiveteen() {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	i := 100
	if i <= len(runes)-1 {
		justString := runes[:100]
		fmt.Println("String:", string(justString))

	} else {
		fmt.Println("Граница слайса больше чем кол-во символов в строке")
	}
}

func createHugeString(value int) string {
	var str string
	for i := 0; i < value; i++ {
		if i%2 == 0 {
			str += "A"
		} else {
			str += "B"
		}
	}
	return str
}
