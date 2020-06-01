package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 4, 5, 6} //array

	for i, numero := range numeros { // i=indice, numero=array[i]
		fmt.Println(i, numero)
	}
	fmt.Println("")

	for _, num := range numeros {
		fmt.Println(num)
	}

}
