package main

import "fmt"

const pi = 3.1415

func main()  {
	printCircleArea(5)
}

func printCircleArea(radius int)  {
	if radius <= 0 {
		fmt.Println("Радиус круга не может быть отрицательным!")
		return
	}
	fmt.Printf("Радиус круга равен %d см\n", radius)
	fmt.Println("Формула для расчета площади круга: S=pi*r*r")

	circleArea := calculateCircleArea(radius)
	fmt.Printf("Площадь круга %.2f см. кв. \n", circleArea)
}
func calculateCircleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi
}