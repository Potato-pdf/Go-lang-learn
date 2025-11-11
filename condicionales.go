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

	//BUcle tupo while
	n := 0
	for n < 3 {
		fmt.Printf("Iteracion : %d\n", n)
		n++
	}

	//bucle infinito
	for {
		n ++
		if n ==5 {
			continue
		}
		fmt.Printf("n en bucle: %d\n", n)
			if n >= 7 {
		break
	}
	}

	//Range
	slice := []string{"2","4","5"}
	for index, value := range slice {
		fmt.Printf("Indice: %d, valor : %s\n", index, value)
	}

	//switch
	valor := 2
	switch valor{
	case 1:
		fmt.Println("Es 1 ")
	
	case 2:
		fmt.Println("Es 2")
	default:
		fmt.Println("No es ni uno ni 2")
	}
}