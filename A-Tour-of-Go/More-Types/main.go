package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}
	p3 = &Vertex{1, 2}
)

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p2 := &v
	p2.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p3, v2, v3)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[2:6]
	fmt.Println(s)

	s[0] = 100
	fmt.Println(primes) // 2, 100, 5, 7, 11, 13

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{7, false},
	}

	fmt.Println(s2)

	fmt.Printf("len=%d cap = %d", len(s), cap(s))

	var emptySlice []int
	if emptySlice == nil {
		fmt.Println("nil!")
	}

	s3 := append(emptySlice, 3, 5)
	fmt.Println(s3)

	for i, v := range s3 {
		// i = index
		// v = value
		fmt.Println(i, v)
	}

	for _, v := range s3 {
		fmt.Println(v)
	}

	var m = map[string]Vertex{
		"Google": Vertex{
			37, -122,
		},
	}

	fmt.Println(m)
	hoge := func() string {
		return "hoge"
	}
	fmt.Println(hoge())

	pos, neg := adder(), adder()
	// sum is saved
	fmt.Println(pos(2), neg(4)) // 2, 4
	fmt.Println(pos(2), neg(4)) // 4, 8
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
