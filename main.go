package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)

func main() {



}






func Shop() {

	fmt.Println("::::::::::::::::::::::::::::::::::::::")
	fmt.Println("::  [1]    ---------------    BUY   ::")
	fmt.Println("::  [2]    ---------------    View  ::")
	fmt.Println("::  [3]    ---------------    Exit  ::")
	fmt.Println("::::::::::::::::::::::::::::::::::::::")

	// Create a new scanner object
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Choose: ")
	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 10, 64 )

	if(input == 1) {
		fmt.Println("BUY")
	} else if (input == 2) {
		fmt.Println("View")
	} else if (input == 3) {
		fmt.Println("Exit")
	} else {
		fmt.Println("Please Choose the right command")
	}



}
















	// // var (Variable Name) (Variable Type) = (Variable Assignment)
	// var firstName string = "Mark Vergel"

	// // Pwede
	// var age = 21
	// var height = 168

	// // Shortcut Declaration
	// lastName 	:= "Banguis"
	// address 	:= "Surigao City"
	// school 		:= "Surigao State College of Technology"
	// company 	:= "Cauld and Clark"
	// weight 		:= 54.305345345

	// fmt.Println("\nFullname: " + firstName + " " + lastName)

	// fmt.Println("Works at: " + company)

	// fmt.Printf("%s %d\n", "Age:", age)

	// fmt.Printf("%s %d\n", "Height:", height)

	// fmt.Printf("%s %.2f\n", "Weight:", math.Floor(weight*100)/100)

	// fmt.Println("Current Adress: " + address)

	// fmt.Println("University: " + school)