package main

import ("fmt")

func one() {
	x := 25 //type inference
	a := &x //memory address where x is being stored

	// & symbol points to a memory address
	// * symbol reads to the memory address

	//overrides the value
	*a = 5
	fmt.Println(*a)
	fmt.Println(x) //value of x is now changed since it was changed through the memory address

	*a = *a * *a
	fmt.Println(x)
}

func main() {
	one()
}
