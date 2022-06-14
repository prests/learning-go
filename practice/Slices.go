package main

import (
	"fmt"
	"sort"
)

func main() {
	//Array
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	//Slice
	var colorsSlice = []string{"Orange", "Pink", "Yellow"}
	colorsSlice = append(colorsSlice, "Purple")
	fmt.Println(colorsSlice)

	colorsSlice = append(colorsSlice[1:len(colorsSlice)])
	fmt.Println(colorsSlice)

	colorsSlice = append(colorsSlice[:len(colorsSlice)-1])
	fmt.Println(colorsSlice)

	//remove capacity to add more
	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 6
	numbers[2] = 34
	numbers[3] = 14
	numbers[4] = 34
	fmt.Println((numbers))
	sort.Ints(numbers)
	fmt.Println(numbers)
}
