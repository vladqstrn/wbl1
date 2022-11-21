package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.
func Fourteen() {
	valueInt := 777
	valueStr := "wb"
	valueBool := true
	valueChan := make(chan int)
	GetType(valueInt, valueStr, valueBool, valueChan)
}

func GetType(valueInt, valueStr, valueBool, valueChan interface{}) {
	fmt.Printf("Type: %T\n", valueInt)
	fmt.Printf("Type: %T\n", valueStr)
	fmt.Printf("Type: %T\n", valueBool)
	fmt.Printf("Type: %T\n", valueChan)
}
