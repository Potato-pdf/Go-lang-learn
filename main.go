package main

import "fmt"


func main(){
	suma:= Suma(3,4)
	fmt.Println(suma)

	cociente, error:=  Division(9,0)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(cociente)

	VaribleDeArgumntos()
}
