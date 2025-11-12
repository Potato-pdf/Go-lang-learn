package domain

import (
	"fmt"
	"learning-go/domain/enities"
)
func App(){
	persona:= enities.Persona{
		Nombre: "Juan",
		Edad:   30,
	}
	elemento:= enities.TipoELemento{
		Nombre: "Agua",
	}
	persona.CumplirAnios()
	saludo:= persona.Saludar()
	fmt.Println(saludo)
	fmt.Println("Edad:", persona.Edad)
	fmt.Println("Elemento:", elemento.Nombre)
}