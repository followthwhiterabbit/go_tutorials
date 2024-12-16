package main 

import "fmt"


// function that takes two integers and returns their sum 

func add(a int, b int ) int {
	return a + b; 

}

func main() {
	/*
	// variable declaration 
	var name string = "Alice"
	var age int = 25

	// short hand declaration
	city := "Istanbul"

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	*/ 

	result := add(5, 7)
	fmt.Println("Sum:", result)


	}
