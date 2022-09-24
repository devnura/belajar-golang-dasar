package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

func main() {
	for i := 0; i < 3; i++ {
		sayHello()
	}
}
