package main

import (
	"fmt"
	"strconv"
)

//declaracion de variables globales
var numero int
var texto string
var status bool

func main() {
	//declara variables locales
	//var numero2, numero3, numero4 int

	//declara variables locales
	numero2, numero3, numero4, texto, status := 2, 5, 67, "valor texto", true

	var moneda int64 = 5
	//convierte de int64 a int generico
	numero2 = int(moneda)
	//convierte de int64 a string
	texto = fmt.Sprintf("%d", moneda)
	//convierte de int64 a string
	texto = strconv.Itoa(int(moneda))

	fmt.Println("numero2 ", numero2, " numero3 ", numero3, " numero4", numero4, "texto ", texto, " status ", status)

	//invocar funcion
	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
