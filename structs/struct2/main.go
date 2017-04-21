package main

import "fmt"

type PointPrinter interface {
	printPoint()
}

type PointB struct {
	p1, p2 int
}

type Point struct {
	x, y int
}

func (p *PointB) printPoint() {
	fmt.Printf("method adress=%v, a=%d, b=%d\n", &p, p.p1, p.p2)
}

func (p *Point) printPoint() {
	fmt.Printf("method adress=%v, a=%d, b=%d\n", &p, p.x, p.y)
}

func createPoint(a, b int) Point {
	return Point{x: a, y: b}
}

func (p *Point) setPoint(a, b int) {
	p.x = a
	p.y = b
}

func printPoint(p PointPrinter) {
	p.printPoint()
}

func updatePoint(p *Point) {
	p.x *= 2
	p.y *= 2
	println(&p)
}

func main() {
	p := createPoint(1, 2)
	printPoint(&p)

	p.setPoint(3, 4)
	printPoint(&p)

	updatePoint(&p)
	printPoint(&p)

	p2 := PointB{32, 33}
	printPoint(&p2)
}
