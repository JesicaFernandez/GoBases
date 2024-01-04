package main

import "fmt"

func main()  {
	
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 38} 
	var cantEmploye = 0
	fmt.Println(employees["Benjamin"])

	for _, employee := range employees {
		
		if employee > 21 {
			cantEmploye++ 
		}
	}

	fmt.Printf("Hay %d empleados mayores a 21 años\n", cantEmploye)
	
	employees["Federico"] = 25
	fmt.Println("Mapa después de agregar a Federico:", employees)

	delete(employees, "Pedro")
	fmt.Println("Mapa después de eliminar a Pedro:", employees)
	
}