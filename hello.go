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

func addition (a float64, b float64) float64 { //variables and return type
	return a+b+rando
}

// A constant value
const rando float64 = 1

// note that that for each return type must be defined, even if they are all the same.
func dblShow (word string, sentence string) (string, string) {
	return word, sentence
}

func main() {
	sayHello()
	randomNumber()
	fmt.Println("While we are at it, the square root of 16 is.., ", math.Sqrt(16))

	// 1. decalre varibables separately
	/*
	var x float64 = 10
	var y float64 = 5
	*/

	// 2. declare variables together in one command
	var x,y float64 = 10, 5

	// 3. since types in golang don't change, you don't need to initiate variables with type
	// NOTE: the variables are determined (if without a type) based on the value. in the case, x&y are int
	//x, y := 10, 5


	//Conversion of types
	//var x float32 = 10
	//var y float64 = float64(x)

	//OR

	//t := x //t will be type x (type inference)

	fmt.Println(dblShow("hello","dd"))
	fmt.Println(addition(x, y))
}