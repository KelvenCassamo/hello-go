package basics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func GoStructs() {
	fmt.Println("Go Structs:")
	p := Person{Name: "Kelven", Age: 21}
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
