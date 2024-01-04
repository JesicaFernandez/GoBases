package main

import (
	"math/rand"
	"fmt"
	"time"
)
/*Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual 
se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas..*/
func main()  {
	
	rand.Seed(time.Now().UnixNano())

	promedio := calculatePromedio(5)
	fmt.Println("El promedio es:",promedio)
}

func calculatePromedio(n int) (promedio float32) {

	suma := 0
	for i := 0; i < n; i++ {
		
		num := rand.Intn(11)
		fmt.Println(num, " ")

		if num < 0 {
			fmt.Println("No se pueden introducir notas negativas")
			break
		}
		
		suma += num
		promedio = float32(suma) / float32(n)
	}
	return
}