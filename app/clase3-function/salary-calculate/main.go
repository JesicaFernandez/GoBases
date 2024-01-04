package main

import (
	"math/rand"
	"fmt"
	"time"
)
/* Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la 
categoría.

1. Categoría C, su salario es de $1.000 por hora.
2. Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
3. Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva 
su salario.*/
func main()  {
	
	rand.Seed(time.Now().UnixNano())

	salary := calculateSalary(60, "A")
	fmt.Println("Su salario es de:",salary)
}

func calculateSalary(min int, category string) (salary float32) {

	switch category {
	case "C":
		salary = float32(min/60) * 1000
	case "B":
		salary = (float32(min/60) * 1500) + (float32(min/60) * 1500 * 0.2)	
	case "A":
		salary = (float32(min/60) * 3000) + (float32(min/60) * 3000 * 0.5)
	default:
		fmt.Println("La categoría no es válida")
	}
	return	 
}