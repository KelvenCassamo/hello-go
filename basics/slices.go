package basics

import "fmt"

func GoSlices() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Go Slices:")
	fmt.Println(slice)
}
