package main

import "fmt"

func main()  {
	
	cantLetras("Computer")
	printLentras("Computer")
}

func cantLetras(palabra string){
	fmt.Println(len(palabra))
}

func printLentras(palabra string){

	for _, caracter := range palabra {
        fmt.Printf("%c ", caracter)
    }
	fmt.Println()
}