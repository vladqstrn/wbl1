package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type OpenDoorService interface {
	OpenHomeDoor()
}

type Home struct {
}

func (h *Home) OpenDoorService() {
	fmt.Println("Дверь в дом открыта!")
}

type Car struct {
}

func (c *Car) OpenDoor() {
	fmt.Println("Дверь машины открыта!")
}

type CarAdapter struct {
	Car *Car
}

func (adapter *CarAdapter) OpenHomeDoor() {
	adapter.Car.OpenDoor()
}

func TwentyOne() {
	h := &Home{}
	h.OpenDoorService()
	c := &Car{}
	ca := &CarAdapter{c}
	ca.Car.OpenDoor()

}
