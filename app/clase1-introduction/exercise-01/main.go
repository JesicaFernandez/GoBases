package main

import "fmt"

func main()  {
	
	//Funciones
	printNameAndAddress("Jesica", "Avenida Lujan")
}

func printNameAndAddress(nombre, direccion string) {

	fmt.Println("Nombre: ", nombre)
	fmt.Println("Dirección: ", direccion)
}