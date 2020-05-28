package main

import (
	"fmt"
)

func main() {

	var i int // Declare
	// i = 42    // Define

	fmt.Println(i)

	// Declare and intialize
	var name string = "Imran"
	fmt.Println(name)

	// Implicit initialization
	isAlive := true
	fmt.Println(isAlive)

	// Complex numbers: Real and Imaginary
	comp := complex(1, 2)
	fmt.Println(comp)

	// Unwrap complex type
	r, x := real(comp), imag(comp)
	fmt.Println(r)
	fmt.Println(x)

	// Pointers
	var fullName *string = new(string)
	*fullName = "Imran Aftab Rana" // De-reference operation to set the value
	fmt.Println(*fullName)         // De-reference operator to get the value

	// Address of Operator
	age := 27
	agePtr := &age
	fmt.Println(age, agePtr, *agePtr)
	*agePtr = 28
	fmt.Println(age, agePtr, *agePtr)
	*(&age) = 29
	fmt.Println(age, agePtr, *agePtr)
	salary := 1000000
	agePtr = &salary
	*agePtr = 10
	fmt.Println(age, agePtr, *agePtr)

	fmt.Println(1 + 2.3)

	const c = 1
	fmt.Println(c + 2.3)

	// Multiple constants and Iota
	const (
		q = iota
		w = iota
		e = iota + 2
		t
		y
		o = "ok"
		p
	)

	fmt.Printf(
		"q = %d , w= %d , e= %d , t= %d , y= %d, o = %s, p = %s \n", q, w, e, t, y, o, p)

	const b = iota // iota resets itself here
	fmt.Println("b = ", b)
}
