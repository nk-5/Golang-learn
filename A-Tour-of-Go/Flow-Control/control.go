package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x)
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v, lim)
	}
	return lim // not return v
}

func f() int {
	return 1
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// optional init and end process
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux. ")
	default:
		// freebsd, window
		fmt.Printf("%s", os)
	}

	switch i := 0; i {
	case 0:
	case f(): // not called
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("hello")

	// defer call when function is return
	// multiple defer order call LIFO
	defer fmt.Println("fuga")
	defer fmt.Println("keigo is cool")
	defer fmt.Println("hoge")

	fmt.Println("world")

	// hello
	// world
	// hoge
	// keigo is cool
	// fuga
}
