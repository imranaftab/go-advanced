package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/imranaftab/go-advanced/controllers"
	"github.com/imranaftab/go-advanced/models"
)

func main() {

	// Setup routing for web service
	controllers.RegisterController()

	// Start the webserver to listen and respond to incoming http requests
	http.ListenAndServe(":1234", nil)

	// Primitive Types
	primitiveTypes()

	// Collections
	collections()

	// Structs
	structs()

	// Utilize a struct defined in a module
	user := models.Person{Name: "Imran", Age: 28, IsAlive: true}
	fmt.Println(user)

	// Normal Function with params
	simpleFunction(1)

	// Normal Function with return
	isSuccessful := funcMultiParamReturn("Imran Aftab", "Rana", 28)
	fmt.Println(isSuccessful)

	// Normal function with multiple parameters of same types and multiple return variables
	isSuccess, value := functWithMutipleReturns()
	fmt.Println(isSuccess, value)

	// Writeable variable (_) usage with Normal function having multiple parameters of same types and multiple return variables.
	_, resultValue := functWithMutipleReturns() // Here we use _ (writeable variable), because we plan to use only one return value and ignore the other
	fmt.Println(resultValue)

	_ = funcReturnNill()

	err := funcReturnError()
	fmt.Println(err)

}

func funcReturnError() error {
	return errors.New("Something went wrong")
}

func funcReturnNill() error {
	return nil
}

func functWithMutipleReturns() (bool, int) {
	return true, 10
}

func funcMultiParamReturn(firstName, lastName string, age int) bool {
	fmt.Printf("Name: %s %s \n", firstName, lastName)
	fmt.Printf("Age: %d\n", age)

	return true
}

func simpleFunction(count int) {
	fmt.Println("Simple function", count)
}

func primitiveTypes() {
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

	// variable blocks
	var (
		fullUsername = "Imran"
		fullAge      = 28
	)

	fmt.Println(fullUsername, fullAge)

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

func collections() {
	// Arrays explicit initialization
	var fruits [4]string
	fruits[0] = "I"
	fruits[1] = "Love you"
	fruits[2] = "Imran"

	fmt.Println(fruits)

	// Array implicit initialization
	nos := [3]int{1, 3, 4}
	fmt.Println(nos)

	// Slices intialization with arrays
	slices := nos[:]
	nos[0] = 2
	println(slices[0], nos[0])

	// Slices implicit intialization
	impSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(impSlice)

	// Append data in slices
	impSlice = append(impSlice, 1)
	fmt.Println(impSlice)

	// Maps
	person := map[string]int{"age": 28}
	fmt.Println(person)
	fmt.Println(person["age"])

	person["age"] = 30
	fmt.Println(person["age"])

	// Add/Updata key from map
	person["name"] = 22
	fmt.Println(person["name"])

	// Delete key from map
	fmt.Println(person)
	delete(person, "age")
	fmt.Println(person)

}

func structs() {

	// Declare a struct - Only scoped to this function, can be moved to package scope
	type person struct {
		name    string
		age     int
		isAlive bool
	}

	// Explicit initialization of struct
	var imran person
	fmt.Println(imran)
	imran.name = "Imran"
	imran.age = 28
	imran.isAlive = true
	fmt.Println(imran)

	// Implicit initialization of struct
	imranaftab := person{
		name:    "Imran Aftab Rana",
		age:     28,
		isAlive: true,
	}
	fmt.Println(imranaftab)
}
