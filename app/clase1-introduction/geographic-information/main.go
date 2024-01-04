package main

import "fmt"

func main()  {
	
	var latitud, longitud float64 = 39.3, 43.2
	var cityName string = "Santa Fe"

	fmt.Printf("La ciudad de %s tiene una latitud de %.2f y una longitud de %.2f\n", cityName, latitud, longitud)
}
