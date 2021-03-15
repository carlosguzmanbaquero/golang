package main

import "fmt"

func main() {
	inicial()
	matriz()
	arreglosDinamicos()
	variante1()
	variante2()
	variante3()
	variante4()
}

func inicial() {
	var elementos [10]int
	elementos[0] = 1
	elementos[5] = 15

	fmt.Println(elementos)
}

func matriz() {
	var matriz [5][7]int
	matriz[3][5] = 1
	fmt.Println(matriz)

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			fmt.Print(matriz[i][j])
		}
		fmt.Println("")

	}
}

//Slice
func arreglosDinamicos() {
	arreglo := []int{2, 5, 4}
	fmt.Println(arreglo)
	matriz := [][]int{{2, 5, 4}, {2, 5, 4}}
	fmt.Println(matriz)
}

func variante1() {
	elementos := [5]int{1, 2, 3, 4, 5}
	fmt.Println(elementos)

	for i := 0; i < len(elementos); i++ {
		fmt.Println(elementos[i])
	}
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	//Slice
	porcion := elementos[2:4]
	fmt.Println(porcion)
	porcion = append(porcion, 8)
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))

}
