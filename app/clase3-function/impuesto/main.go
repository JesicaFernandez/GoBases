package main

import "fmt"
/*Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es 
necesario crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de 150.000 se le descontará 
además un 10 % (27% en total).*/
func main()  {
	
	sueldo := calculateImpuesto(1000000)
	fmt.Println("Su sueldo con los descuentos es:",sueldo)
}

func calculateImpuesto(sueldo float32) float32 {

	if sueldo >= 150000 {
		sueldo = sueldo - (sueldo * 0.27)
	} else if sueldo >= 50000 {
		sueldo = sueldo - (sueldo * 0.17)	
	} 
	return sueldo
}