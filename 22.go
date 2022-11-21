package main

import (
	"fmt"
	"math/big"
	"os"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
// значение которых > 2^20.
func TwentyTwo() {

	var value1 string
	var value2 string
	var choice string
	fmt.Println("Введите первое значение")
	fmt.Fscan(os.Stdin, &value1)
	fmt.Println("Введите второе значение")
	fmt.Fscan(os.Stdin, &value2)
	a := new(big.Int)
	a.SetString(value1, 10)
	b := new(big.Int)
	b.SetString(value2, 10)
	fmt.Println("Выберите действие:  +, *, -, /")
	fmt.Fscan(os.Stdin, &choice)

	buff := new(big.Int)

	switch choice {
	case "+":
		fmt.Println(buff.Add(a, b))
	case "*":
		fmt.Println(buff.Mul(a, b))
	case "-":
		fmt.Println(buff.Sub(a, b))
	case "/":
		fmt.Println(buff.Div(a, b))
	}

}
