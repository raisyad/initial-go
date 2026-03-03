package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	Pi     complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func main() {
	var i int
	fmt.Println(i, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Pi, Pi)

	sayHello()
	fmt.Println(add(1, 2))
	fmt.Println(split(17))
	fmt.Println(swap("hello", "world"))
}

func sayHello() {
	fmt.Println("Hello this is practice to create function")
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}
