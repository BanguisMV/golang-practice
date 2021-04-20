package main

import (
	"fmt"
)

func main() {

	// var (Variable Name) (Variable Type) = (Variable Assignment)
	var firstName string = "Mark Vergel"

	// Shortcut Declaration
	lastName := "Banguis"
	age := 21
	address := "Surigao City"
	school := "Surigao State College of Technology"
	company := "Cauld and Clark"
	weight := 54.50
	height := 168

	fmt.Println("\nFullname: " + firstName + " " + lastName)

	fmt.Println("Works at: " + company)

	fmt.Printf("%s %d\n", "Age:", age)

	fmt.Printf("%s %d\n", "Height:", height)

	fmt.Printf("%s %f\n", "Weight:", weight)

	fmt.Println("Current Adress: " + address)

	fmt.Println("University: " + school)

}
