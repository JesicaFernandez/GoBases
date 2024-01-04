package main

import "fmt"

func main()  {
	
	otorgarPrestamo(21, true, 2, 90000.90)
}

func otorgarPrestamo(age int, isEmpleado bool, antiquity int, salary float32){

	if age > 22 && isEmpleado && antiquity > 1 {
		if salary > 100000 {
			// Caso 1: Cliente elegible para préstamo sin interés
			fmt.Println("¡Felicidades! Eres elegible para un préstamo sin interés.")
		} else {
			// Caso 2: Cliente elegible para préstamo con interés
			fmt.Println("Eres elegible para un préstamo con interés.")
		}
	} else {
		// Caso 3: Cliente no elegible para préstamo
		fmt.Println("Lo siento, no cumples con los requisitos para obtener un préstamo.")
	}
}