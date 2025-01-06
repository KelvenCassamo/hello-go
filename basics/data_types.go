package basics

import "fmt"

func GoDataTypes() {
	var i int = 42
	var f float64 = 3.14
	var f32 float32 = 3.14
	var s string = "GoLang"
	var b bool = true

	fmt.Println("Go Data Types:")
	fmt.Printf("int: %d, float64: %.2f, float32: %.2f, string: %s, bool: %t\n", i, f, f32, s, b)
}
