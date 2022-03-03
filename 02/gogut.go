package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square foot of 4 is", math.Sqrt(4))
}

func main() {
	foo()
	fmt.Println("A number from 0-100", rand.Intn(101))
}