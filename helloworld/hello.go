package main

import "fmt"

func main() {
    fmt.Printf("hello world, sum=%d\n", add(5,4))
}

func add(x int, y int) int {
    return x + y
}
