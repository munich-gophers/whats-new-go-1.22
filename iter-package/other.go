package main

import (
	"fmt"
)

// Person struct to hold individual data
type Person struct {
	Name string
	Age  int
}

// PersonIterator function to iterate over a slice of Person
func PersonIterator(people []Person) func(func(int, Person) bool) {
	return func(yield func(int, Person) bool) {
		for i, p := range people {
			if !yield(i, p) { // If yield returns false, stop iteration
				return
			}
		}
	}
}

func main() {
	// Sample data
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Using the iterator to print each person's details
	for i, person := range PersonIterator(people) {
		fmt.Printf("Person %d: Name: %s, Age: %d\n", i, person.Name, person.Age)
	}
}
