package main

import (
	"errors"
	"fmt"
)

func Suma(a int, b int ) int {	
	return a+b
}

func Division(a,b float64)(error, float64){
	if b ==0{
		fmt.Errorf("No se puede dividir entre 0")
		return errors.New("No se puede dividir por 0"), 0
	}
	cociente := a/b
	return nil, cociente
}