package main

import (
	"fmt"
	"time"
)

//
func tarea(id int, canal chan <- string){
	time.Sleep(time.Duration(id)*100*time.Millisecond)
	canal<- fmt.Sprintf("Tarea %d terminada \n", id)
}

func main(){
	canal:= make(chan string, 2)
	for i:= range 4 {
		go tarea(i, canal)

	}
	timeout := time.After(1 * time.Second)
	for range 4 {
		select {
		case msj := <-canal:
			fmt.Printf(msj)
		
		case <- timeout:
		fmt.Println("Tiempo de espera agotado")
	}
	}
}