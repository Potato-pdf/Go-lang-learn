package main

import "fmt"


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

}
