package main
import "fmt"

func Tarea(id int, canal chan <-string){
	canal <- fmt.Sprintf("Tarea %d completada", id)
}
func main() {
	canal := make (chan string)
	//Solo puedo enviar strings por este canal, el canal es una conexion entre go rutines
	for i:= range 3{
		go Tarea(i,canal)
		//go rutine : una tarea que se ejecuta de manera concurrente
	}
	for i := 0 ; i<3;i++{
		mensaje := <- canal
		fmt.Println(mensaje)
	}
}