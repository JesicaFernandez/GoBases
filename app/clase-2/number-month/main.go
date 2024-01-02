package main

import "fmt"

func main()  {
	
	var numberMonth int = 18

	switch {
	case numberMonth == 1:
		fmt.Println("Juanary")
	case numberMonth == 2:
		fmt.Println("February")
	case numberMonth == 3:
		fmt.Println("March")
	case numberMonth == 4:
		fmt.Println("April")
	case numberMonth == 5:
		fmt.Println("May")
	case numberMonth == 6 :
		fmt.Println("June")
	case numberMonth == 7:
		fmt.Println("July")
	case numberMonth == 8:
		fmt.Println("August")
	case numberMonth == 9:
		fmt.Println("September")
	case numberMonth == 10:
		fmt.Println("October")
	case numberMonth == 11:
		fmt.Println("November")
	case numberMonth == 12:
		fmt.Println("December")
	default:
		fmt.Println("Número del mes ingresado incorrecto!")
	}
	
}
