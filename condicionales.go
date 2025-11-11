package main

import "fmt"

func Condicional(){

	//Defer
	defer fmt.Println("Fin") //Se eejecuta al final, una vez que ya se ejecutara todo lo relacionado a la funcion en la que se encuentra

	//If
	edad :=20
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	}else {
		fmt.Println("Eres menor de edad")
	}
	//Asertive/negative programing
	if edad < 18{
		fmt.Println("Eres menor de edad")
	}
	
	//Bucle classico
	for i := 0; i < 5 ; i ++{
		fmt.Printf("Iteracion: %d\nk", i)
	}

}