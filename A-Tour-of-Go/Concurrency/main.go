package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to C
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	fmt.Println("hoge")
	go sum(s[len(s)/2:], c)
	fmt.Println("fuga")
	x, y := <-c, <-c // receive from c
	fmt.Println("keigo")

	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch2 := make(chan int, 10)
	go fibonacci(cap(ch2), ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	ch3 := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go ch3.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(ch3.Value("somekey"))
}
