package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("hola")
	//Numeros
	var entero int = 10 
	entero1 := 10 
	decimal := 3.24
	suma := entero + int(decimal) + entero1
	fmt.Println(suma)
	//String
	mensaje := "hola, tu edad es"
	edad := strconv.Itoa(suma)
	concatenado := mensaje + edad 
	fmt.Println(concatenado)

	//Booleans
	esVerdadero:= true
	fmt.Println(esVerdadero)

	//Arrays
	arrayFijo := [2]int{1,2} //el primer [] dice cuantos datos va a tener el array
	sliceVariable :=  []int{3,2,1,23,} // con slice se puede escribir sin limitante en el arreglo
	fmt.Println(sliceVariable, arrayFijo)
	sliceVariable = append(sliceVariable, 55) //agrega datos al array 
	fmt.Println(sliceVariable)

	//Maps
	diccionario := map[string]int{ //En los map el [] indica el tipo de la clave, mientras que en este caso el int es el tipo del valor
		"clave1":1,
		"clave2":2,
	}
	fmt.Println(diccionario)

	//Structs
	type Person struct { //Si inicia con mayuscula significa que se va a exportar, si esta en minuscula, esolo espara ese archivo 
		Nombre string
		Edad int
	}
	persona := Person{Nombre : "Cesar", Edad : 23}
	edad_to_string:= strconv.Itoa(persona.Edad)
	fmt.Println(
		"tu nombre es " + persona.Nombre, "y tu edad es " + edad_to_string, 
	);

	// convertir tios automaticamente con Go
	fmt.Printf("Tu nombre es %s y tu edad es %d\n", persona.Nombre, persona.Edad)

//Tipos de datos
//int, int64, int32 | entero | siempre se use int a menos que se requiera control sobre el tamaño
//float32, float64 | se usa para representar valores numericos reales(decimales) | 32 o 64 va asociado a la arquitectura del sistema
//uint, uint16, uilt32, uint64 | entero sin signo (positivos unicamente)
//bool |true/false | flag o condicional
//string |"cadena de caracteres"
//byte | === unit8 | se usa para datos binarios
//rune | === int32 | se usa para representar un sol ocaracter que ocupa más de 1 byte
//complex6, complex128 | cuando tiene una parte real y una imaginaria 

}
