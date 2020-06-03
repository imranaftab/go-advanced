package models

import (
	"errors"
	"fmt"
)

// Person Person struct with person data
type Person struct {
	ID      int
	Name    string
	Age     int
	IsAlive bool
}

var (
	people []*Person
	nextID = 1
)

//====================================//
// In-Memory CRUD operations with GO  //
//====================================//

// AddPeople Add people to the list of pointers
func AddPeople(person Person) (Person, error) {
	if person.ID != 0 {
		return Person{}, errors.New("Person must have ID equal to 0")
	}
	person.ID = nextID
	nextID++
	people = append(people, &person)

	return person, nil
}

// GetPeople Gets pointer to the list of persons
func GetPeople() []*Person {
	return people
}

// GetPeopleByID Finds the person with the given id
func GetPeopleByID(id int) (Person, error) {
	for _, person := range people {
		if id == person.ID {
			return *person, nil
		}
	}

	return Person{}, fmt.Errorf("Unable to find person with id %d", id)
}

// UpdatePeople Updates the user with the given id
func UpdatePeople(per Person) (Person, error) {
	for i, person := range people {
		if per.ID == person.ID {
			people[i] = &per
			return *person, nil
		}
	}

	return Person{}, fmt.Errorf("Unable to update person with id %d. Person not found", per.ID)
}

// DeletePeople Deletes person with the given id
func DeletePeople(id int) (Person, error) {
	for i, person := range people {
		if id == person.ID {
			// here we create new collection by appending a slice of people that does not include the person with given id
			// and appends list of the rest of the people by using spread operator (...) on collection
			people = append(people[:i], people[i+1:]...)
			return *person, nil
		}
	}

	return Person{}, fmt.Errorf("Unable to delete person with id %d. Person not found", id)
}
