package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("random number is", rand.Intn(100))

	fmt.Printf("Now you have %g problems. \n", math.Sqrt(100))
}
