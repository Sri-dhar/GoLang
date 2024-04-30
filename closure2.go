package main

import "fmt"

// generateFloatArrayGenerator is a function that returns another function
// The returned function generates and returns a float array of size n
func generateFloatArrayGenerator() func(int) []float64 {
	// This inner function captures the 'size' variable from its outer scope
	return func(size int) []float64 {
		// Initialize a float array of size 'size'
		arr := make([]float64, size)

		// Fill the array with some arbitrary values (e.g., 1.0, 2.0, ..., n)
		for i := 0; i < size; i++ {
			arr[i] = float64(i + 1)
		}

		// Return the generated float array
		return arr
	}
}

func main() {
	// Create a float array generator
	generator := generateFloatArrayGenerator()

	// Generate and print a float array of size 5
	arr1 := generator(5)
	fmt.Println("Array 1:", arr1)

	// Generate and print a float array of size 3
	arr2 := generator(3)
	fmt.Println("Array 2:", arr2)
}
