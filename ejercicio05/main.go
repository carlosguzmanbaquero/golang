package main

import (
	"fmt"
)

//declaracion de variables globales

func main() {
	/*i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	numero := 0
	for {
		fmt.Println("Continuar")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n valor de i: %d", i)
		if i == 5 {
			fmt.Printf("\n multiplicamos por 2 \n")
			i = i * 2
			continue
		}
		fmt.Printf("\n lee esta linea \n")
		i++
	}
	*/

	var i int = 0

RUTINA:
	for i < 10 {
		fmt.Printf("\n valor de i: %d", i)
		if i == 4 {

			i = i + 2
			fmt.Println("voy a RUTINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("valor de i: %d\n", i)
		i++
	}

}
