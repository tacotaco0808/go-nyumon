package main

import "fmt"

func main() {
	rectangle := Rectangle{
		Width:  30,
		Height: 15.5,
	}

	circle := Circle{
		Radius: 10,
	}

	printArea(rectangle)
	printArea(circle)
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func printArea(s Shape) {
	fmt.Println("面積: ", s.Area(), "㎠")
}
