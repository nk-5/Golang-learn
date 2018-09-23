package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) int {
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

var c, python, java bool
var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	s      string // "" empty value
)

const Pi = 3.14

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(3))

	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k, c, python, java)

	// convert type
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println("Happy", Pi, "Day")
}
