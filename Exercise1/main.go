package main

import "fmt"

var x = 40

const y int = 41

func main() {
	fmt.Printf("the value of x is %v and the type of x is %T", x, x)
	fmt.Printf("the value of y is %v and the type of y is %T", y, y)

	z := 42
	fmt.Printf("the value of z is %v and the tzpe of z is %T", z, z)
}
