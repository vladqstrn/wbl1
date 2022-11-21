package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде
//структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func CreatePoint(valueX, valueY float64) *Point {
	return &Point{
		x: valueX,
		y: valueY,
	}
}

func CalcDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Abs((math.Pow(p2.x, 2) - math.Pow(p1.x, 2)) + (math.Pow(p2.y, 2) - math.Pow(p1.y, 2))))
}

func TwentyFour() {
	pOne := Point{22, 5}
	pTwo := Point{17, 21}
	res := CalcDistance(pOne, pTwo)
	fmt.Println(res)

}
