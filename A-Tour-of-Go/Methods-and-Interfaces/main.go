package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type MyFloat float64

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	var a Abser
	a = f
	a = &v
	// a = v  compile error, Abs define *Vertex

	fmt.Println(a.Abs())

	var i I
	fmt.Println(i) // nil
	// i.M() runtime error. because M() method nothing

	var i2 interface{} = "hello"

	s := i2.(string)
	fmt.Println(s)

	s, ok := i2.(string)
	fmt.Println(s, ok)

	f2, ok := i2.(float64)
	fmt.Println(f2, ok)

	// f2 = i2.(float64) panic

	do(21)
	do("hello")
	do(true)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
