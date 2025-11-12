package main

import (
	"fmt"
	"sync"
	"time"
)

func ImprimiArea(f Forma) {
	fmt.Printf("El area es: %.2f \n", f.Area())
}

func main(){
	suma:= Suma(3,4)
	fmt.Println(suma)

	cociente, error:=  Division(9,9)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(cociente)

	VaribleDeArgumntos()

	cont := Closure()
	fmt.Println(cont())
	fmt.Println(cont())
	fmt.Println(cont())
	fmt.Println(cont())

	rect := Rectangle{Ancho: 10, ALto: 69}
	fmt.Println("Area del rectangulo:", rect.Area())

	valor := 10
	Incrementar(&valor)
	fmt.Println(valor)


	puntero := new(int) //puntero iniciado en 0
	fmt.Println(puntero)
	
	c:= Circle{Radio: 5}
	ImprimiArea(c)

	canal := make(chan string)

	go DecirHola(canal)
	IMprimirMensaje(canal)

	canal2 := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			canal2 <- i
		}
		close(canal2)
	}()

	for valor := range canal2 {
		fmt.Println(valor)
	}

	//mutex
	var contador int
	var mu sync.RWMutex

	//Writer
	go func (){ 
		for i := 0 ; i < 5 ; i ++ { 
			mu.Lock() //Bloqueo de recursos mientras se escribe 
			contador ++
			mu.Unlock()
		}
	}()
	//Reader
	for i :=0 ; i < 5 ; i ++ {
		go func (id int)  {
			for j:=0 ; j < 5 ; j ++ {
				mu.RLock()
				fmt.Printf("Goroutine %d: %d\n", id, contador)
				mu.RUnlock()
				time.Sleep(100*time.Millisecond)
			}
		}(i)		
	}
	time.Sleep(2* time.Second)
}

