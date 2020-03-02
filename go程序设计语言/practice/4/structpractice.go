package main

import "fmt"

//Center the coordinate
type Center struct {
	x, y int
}

//Circle common circle
type Circle struct {
	Center
	Radius int
}

//Wheel common wheel
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	wheel := Wheel{
		Circle: Circle{
			Center: Center{
				x: 8,
				y: 9,
			},
			Radius: 7,
		},
		Spokes: 2,
	}

	fmt.Printf("wheel.x:%v, wheel.Circle.Center.x:%v\n", wheel.x, wheel.Circle.Center.x)
	fmt.Printf("wheel.x:%#v, wheel.Circle.Center.x:%#v\n", wheel.x, wheel.Circle.Center.x)
	fmt.Println(wheel)
	fmt.Printf("%#v", wheel)
}
