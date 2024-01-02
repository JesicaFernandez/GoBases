package main

import "fmt"

func main()  {
	//var student1 string = 23 // incorrecto (var student1 string = "23")
	//var grade float64 = "A" // incorrecto (var grade string = "A")
	//var isEnrolled = "yes" // correcto
	//var number_of_subjects int
	//number_of_subjects := 5 // incorrecto (var number_of_subjects int = 5)

	var student1 string = "23"
	var grade string = "A"
	var isEnrolled = "yes"
	var number_of_subjects int = 5

	fmt.Println("Variables:", student1, grade, isEnrolled, number_of_subjects)
}
