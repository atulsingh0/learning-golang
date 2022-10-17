package main

import "fmt"

func main() {

	// Slice is a projection of underlying array
	// Slices are by default referecing the data when creating a copy

	planets := []string{"Murcury", "Venus", "Earth", "Mars", "Jupitor", "Saturn", "Uranus", "Neptune"}

	fmt.Println("Planets are: ", planets)
	fmt.Printf("Planets are: %T type and having length of %v with capacity of %v\n", planets, len(planets), cap(planets))

	// creating a slice
	terrestial := planets[:4]
	fmt.Printf("Terrestial are %v and having length of %v with capacity of %v\n", terrestial, len(terrestial), cap(terrestial))

	terrestial_2 := planets[:4:4] // capping the capacity as well
	fmt.Printf("Terrestial_2 is having capacity of %v\n", cap(terrestial_2))

	// creating a slice

	slice_1 := make([]int, 10)
	slice_2 := make([]int, 0, 10) // length 0, capacity 10

	fmt.Printf("Slice_1 - Length: %v, Capacity: %v\n", len(slice_1), cap(slice_1))
	fmt.Printf("Slice_2 - Length: %v, Capacity: %v\n", len(slice_2), cap(slice_2))

	// If slice length is 0, we need to use append to add the values

	slice_1[0] = 12 // slice length is not 0 for slice_1 hence replacing the index 0 value with 12
	fmt.Printf("Slice_1 - Length: %v, Capacity: %v\n", len(slice_1), cap(slice_1))

	slice_2 = append(slice_2, 2, 3, 4)
	fmt.Printf("Slice_2 - Length: %v, Capacity: %v\n", len(slice_2), cap(slice_2))

	data := []int{4, 5, 6}
	slice_2 = append(slice_2, data...)
	fmt.Printf("Slice_2 - Length: %v, Capacity: %v\n", len(slice_2), cap(slice_2))
	fmt.Println("Slice_2 data: ", slice_2)

	slice_2 = append(slice_2, slice_2...)
	fmt.Printf("Slice_2 - Length: %v, Capacity: %v\n", len(slice_2), cap(slice_2))
	fmt.Println("Slice_2 data: ", slice_2)

}
