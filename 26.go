package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func TwentySix() {
	str := "abcdefF"
	str = strings.ToLower(str)
	fmt.Println(CheckUnique(str))
}

func CheckUnique(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}
