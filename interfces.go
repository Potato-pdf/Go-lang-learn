package main

import "math"

//Las interfaces no se implementan manualmente
//Las interfaces se cumplen
type Rectangle struct{
	Ancho, ALto float64
}

type Forma interface{
	Area() float64
}

type Circle struct{
	Radio float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radio * c.Radio
}
