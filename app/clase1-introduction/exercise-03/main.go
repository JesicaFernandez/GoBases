package main

import "fmt"

func main()  {
	
	var firstName string = "Jesica"// --> Incorrecta (var firstName string)

    var lastName string = "Yamila" // --> Correcta

    var age int // --> Incorrecta (var age int)

   // lastName := 6 // --> Incorrecta (lastName :=6)

    var driver_license = true // --> Incorrecta (var drive_license := true)

    var person_height int = 3// --> Incorrecta (var person_height int)

    childsNumber := 2 // --> Correcta

	fmt.Println("Variables incorrectas:", age, firstName, lastName, driver_license, person_height, childsNumber)

}