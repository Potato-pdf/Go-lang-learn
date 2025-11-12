package main

// raceondition : ocurre cuando dos o mas go rutines intentan acceder y modificar una variable compartida al mismo tiempo
//mutex : es un mecanismo de sincronizacion que permite asegurar que solo una go rutine pueda acceder a una seccion critica del codigo a la vez
import "fmt"
import "sync"

var (
	completadas int
	mu sync.RWMutex
)

func tarea(id int, canal chan <- string){
	mu.Lock()
	completadas ++
	mu.Unlock()
	canal<-fmt.Sprintf("tarea %d completada", id )
}

func main() {
	canal := make(chan string)

	for i := range 5{
		go tarea(i, canal)
	}

	for range 5 {
		fmt.Println(<-canal)
	}
	mu.Lock()
	fmt.Printf("Total tareas completadas: " ,completadas)
	mu.Unlock()
}