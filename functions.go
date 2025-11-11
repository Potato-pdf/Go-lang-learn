package main

import (
	"errors"
	"fmt"
)

func Suma(a int, b int ) int {	
	return a+b
}

func Division(a,b float64)(float64, error){
	if b ==0{
		return 0, errors.New("no se puede dividir por 0")
	}
	cociente := a/b
	return cociente, nil
}

func VaribleDeArgumntos(nombres ...string){
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

func Closure() func() int{
	c := 0
	return func () int {
		c ++ 
		return c
	}
}