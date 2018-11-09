package main

import (
	"fmt"
	"math"
	"math/rand"
)

func sayHello() {
	fmt.Println("Hello world!")
}

func randomNumber() {
	fmt.Println("A random number: ", rand.Intn(100))
}

func main() {
	sayHello()
	randomNumber()
	fmt.Println("While we are at it, the square root of 16 is.., ", math.Sqrt(16))
}