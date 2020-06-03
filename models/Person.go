package models

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

// GetPeople Gets pointer to the list of persons
func GetPeople() []*Person {
	return people
}

// AddPeople Add people to the list of pointers
func AddPeople(person Person) (Person, error) {
	person.ID = nextID
	nextID++
	people = append(people, &person)

	return person, nil
}
