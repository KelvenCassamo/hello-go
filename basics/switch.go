package basics

import "fmt"

func GoSwitch() {
	day := 3
	fmt.Println("Go Switch:")
	switch day {
	case 1:
		fmt.Println("Segunda")
	case 2:
		fmt.Println("Terça")
	default:
		fmt.Println("Outro dia")
	}
}
