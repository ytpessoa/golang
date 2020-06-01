package main

import "fmt"

func main() {

	// map deve ser inicializado !!!
	mapAprovados := make(map[int]string) //map[key]value
	mapAprovados[123] = "Maria"
	mapAprovados[456] = "Pedro"

	fmt.Println(mapAprovados)
	fmt.Println("")

	for cpf, nome := range mapAprovados {
		fmt.Printf("nome: %s , cpf: %d \n", nome, cpf)
	}

	fmt.Println(mapAprovados[123])
	//Maria
	delete(mapAprovados, 123)

	for cpf, nome := range mapAprovados {
		fmt.Printf("nome: %s , cpf: %d \n", nome, cpf)
	}

}
