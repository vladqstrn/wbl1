package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name   string
	Age    int
	Action // Встраивание структуры Action в структуру Human
}

// GetInfo печатает информацию о человеке
func (h *Human) GetInfo() {
	fmt.Println("Name: ", h.Name, " Age: ", h.Age)
}

type Action struct {
	isWorking bool
}

// Work выводит информацию работает ли человек или нет
func (a *Action) Work() {
	fmt.Println("working = ", a.isWorking, " Name")
}

func Emb() {
	a := Human{"Vlad", 18, Action{true}}
	a.GetInfo()
	//a.Action.Work()
	a.Work()
}

// go начинает проверять наличие поля от корня типа к вложенным типам.
// shadowing, collining
