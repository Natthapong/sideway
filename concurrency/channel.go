package main

import "fmt"

type myfunc func(a, b float64) float64

func say(name string, c chan myfunc) {
	fmt.Println("hello", name)

	fn := func(a, b float64) float64 {
		return a * b
	}
	c <- fn
}

// SISD single instruction single data
// SIMD single instruction multi data
// MISD multi  instruction single data
// MIMD multi  instruction multi data

func main() {
	c := make(chan myfunc)

	go say("natthapong", c)

	fmt.Println("wait")
	fn := <-c //blocking by default
	fmt.Println("done!", fn(3, 4))
}
