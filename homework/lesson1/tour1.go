package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("random number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(100))
	fmt.Println(math.Pi)
	fmt.Println(add(17, 43))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
