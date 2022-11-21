package main

import (
	"fmt"
	"os"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в
//1 или 0.

func Eight() {
	fmt.Println("Введите число: ")
	var value int64
	var bit int64
	var change int64
	fmt.Fscan(os.Stdin, &value)
	fmt.Println("Введите bit: ")
	fmt.Fscan(os.Stdin, &bit)
	fmt.Println("На что меняем (0 или 1): ")
	fmt.Fscan(os.Stdin, &change)
	if change == 1 || change == 0 {
		BitChanger(value, bit, change)
		fmt.Printf("%b, число в 2-ом виде\n", value)
		return
	}
	fmt.Println("Неправильный ввод!")

}

func BitChanger(src int64, position int64, value int64) {
	tmp := uint64(src)
	mask := uint64(1 << position)
	switch value {
	case 0:
		fmt.Printf("%b, число в 2-ом виде\n", int64(tmp&^mask))
		fmt.Println(int64(tmp &^ mask))
	default:
		fmt.Printf("%b, число в 2-ом виде\n", int64(tmp|mask))
		fmt.Println(int64(tmp | mask))
	}
}
