package main

import (
	"cmp"
	"fmt"
	"math"
)

func main() {
	var myFloatValue float64 = 20
	var myFloatValue2 float64 = 20

	/*
		Devuelve 0 si son iguales, 1 si x es mayor a y; y -1 si x es menor a y
	*/

	fmt.Printf("Numbers: (%f, %f) - Compare: %d\n", myFloatValue2, myFloatValue, cmp.Compare(myFloatValue, myFloatValue2))

	//En el entorno normal de comparaciones > < ==, no funciona comprar valores NaN. Compare soluciona eso
	nanvalue := math.NaN()
	fmt.Printf("Numbers: (%f, %f) - Compare: %t\n", nanvalue, myFloatValue, cmp.Less(nanvalue, myFloatValue))

}
