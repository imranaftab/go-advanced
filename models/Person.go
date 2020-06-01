package models

type Person struct {
	Name    string
	Age     int
	IsAlive bool
}

var (
	people []*Person
	nextID = 1
)
