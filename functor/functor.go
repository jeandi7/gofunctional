package main

import "fmt"

// Functor is a structure representing a functor.
type Functor struct {
	data []int
}

// NewFunctor creates a new Functor with the given data.
func NewFunctor(data []int) *Functor {
	return &Functor{data: data}
}

// Map applies a transformation function to each element of the Functor's data.
func (f *Functor) Map(transform func(int) int) *Functor {
	result := make([]int, len(f.data))
	for i, value := range f.data {
		result[i] = transform(value)
	}
	return &Functor{data: result}
}

func main() {
	// Create a Functor with some initial data.
	myFunctor := NewFunctor([]int{1, 2, 3, 4, 5})

	// Define a transformation function (multiply each element by 2).
	transformFunc := func(x int) int {
		return x * 2
	}

	// Apply the transformation using Map.
	transformedFunctor := myFunctor.Map(transformFunc)

	// Print the original and transformed data.
	fmt.Println("Original data:", myFunctor.data)
	fmt.Println("Transformed data:", transformedFunctor.data)
}
