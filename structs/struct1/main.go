package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p := new(Point)
	p.x = 5
	p.y = 8
	fmt.Println("x=", p.x, "y=", p.y)
}
