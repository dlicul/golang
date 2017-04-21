package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	var x, y int = 5, 6
	x2, y2 := swap(x, y)
	fmt.Printf("swap of %d, %d is %d, %d\n", x, y, x2, y2)
}
