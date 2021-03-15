package main

import (
	"fmt"
)

//declaracion de variables globales
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 24 + 7 = %d \n ", calculo(24, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Restar 24 - 7 = %d \n ", calculo(24, 7))

	calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Dividir 24 / 7 = %d \n ", calculo(24, 7))

	Operaciones()

	//clousures
	tablaDel := 2
	miTabla := Tabla(tablaDel)

	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}

	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
