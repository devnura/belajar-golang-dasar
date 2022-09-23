package main

import "fmt"

func main() {
	const firstName string = "Jhon"
	const lastName string = "Doe"
	const value = 1000
	fmt.Println(firstName, lastName, value)

	const (
		firstName1 string = "Jon"
		lastName1  string = "Doe"
		value1     int    = 1000
	)
	fmt.Println(firstName1, lastName1, value1)
}
