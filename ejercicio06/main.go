package main

import (
	"fmt"
)

//declaracion de variables globales

func main() {
	//fmt.Println(uno(5))
	//fmt.Println(dos(1))

	//numero, estado := dos(1)

	//fmt.Println(numero)
	//fmt.Println(estado)

	fmt.Println(calculo(1, 3))
	fmt.Println(calculo(1, 3, 4))

	fmt.Println("\ntotal: ", calculo2(1, 3, 4))
}

/*func uno(numero int) int {
	return numero * 2
}*/

//variante de la funcion anterior
func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		//fmt.Println(num, total)
		total = total + num
	}

	return total
}

func calculo2(numero ...int) int {
	total := 0
	for item, num := range numero {
		fmt.Printf("\nitem %d", item)
		total = total + num
	}

	return total
}
