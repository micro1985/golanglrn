package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var c, python, java bool
var i2, j2 int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("random number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(100))
	fmt.Println(math.Pi)
	fmt.Println(add(17, 43))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(27))

	var i int
	fmt.Println(i, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(i2, j2, c2, python2, java2)

	var i3, j3 int = 1, 2
	k3 := 3
	c3, python3, java3 := true, false, "no!"
	fmt.Println(i3, j3, k3, c3, python3, java3)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x1, y1 int = 3, 4
	var f1 float64 = math.Sqrt(float64(x1*x1 + y1*y1))
	var z1 uint = uint(f1)
	fmt.Println(x1, y1, z1)

	v := 45i // change me!
	fmt.Printf("v is of type %T\n", v)

	const World = "Earth"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
