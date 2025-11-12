package main

import (
	"fmt"
	"time"
)

//go -> GoRoutine
//GoRutine -> es un hilo de ejecucion virtual
//Chanell -> comuicacion entre GoRUtines

func DecirHola(canal chan <- string){
	time.Sleep(2 * time.Second)
	canal <- "Hola Mundo"
}

func IMprimirMensaje(canal <- chan string){
	mensaje := <- canal
	fmt.Println(mensaje)
}