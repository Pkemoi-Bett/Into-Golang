package main

import "fmt"

// Declaring constants

const loginToken = "4r2bu6h3z9" //puplic

func main() {

	// Declaring variables using no var style(Note it is not allowed outside the method)
	username := "Kip"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T\n", username)

	// Boolean Variables

	isLoggedIn := true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T\n", isLoggedIn)

	// Integer
	// There may different type of int you can explore in the go documentation

	age := 20
	fmt.Println(age)
	fmt.Printf("Variable is of Type: %T\n", age)

	// float
	// same to int there are different types of float i.e float32 and float64
	score := 22.3
	fmt.Println(score)
	fmt.Printf("Variable is of Type: %T\n", score)

	fmt.Println(loginToken)
	fmt.Printf("Variable is of Type: %T\n", loginToken)

}
